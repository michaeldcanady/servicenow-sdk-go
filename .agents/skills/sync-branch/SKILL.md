---
name: sync-branch
description: >-
  Identifies the base branch of a provided branch (default: current), reconciles local with
  remote copies, and brings the child branch up to date with its base. Resolves the base from
  the caller, the branch's open-PR base, a stored base/upstream hint, or the repo default
  branch. Then checks whether local and remote copies of the parent exist; fast-forwards a
  stale local parent to match remote; and updates the child by fast-forward when it is strictly
  behind, by merge (the generally-preferred type) when it has diverged, and by rebase only for
  a never-pushed branch with explicit user confirmation. Use whenever the user asks to "sync
  X with its base", "bring X up to date with its history branch / main", "update this branch
  with origin/main", "is my branch behind the base", or wants a long-running branch reconciled
  with the branch it was cut from. Not for creating branches (branch), shipping a change
  (commit-push-pr), or resolving issues (check-related-issue).
---

# Sync a branch with its base

`git pull` only knows the upstream *push* target, not the branch a topic was
actually cut from — and a branch can be several levels deep (feature cut
from feature, feature cut from a `release/vX.Y` line). This skill closes
that gap: it names the base branch, gets local and remote copies into
agreement, then moves the child onto the parent's current tip using the
least disruptive update that fits.

The update type ladder, in preference order (merge is generally preferred):

| Child state vs. parent          | Type                          | Command                       |
| ------------------------------- | ----------------------------- | ----------------------------- |
| strictly behind (behind, 0 ahead) | fast-forward (`--ff-only`)  | `git merge --ff-only <src>`   |
| diverged (behind and ahead)     | merge (merge commit)          | `git merge <src>`             |
| diverged + never pushed + user confirms | rebase (plain)        | `git rebase <src>`            |
| ahead of parent, 0 behind       | no-op                         | — report only                 |
| equal                           | no-op                         | — report only                 |

Never `--force`, `--force-with-lease`, or rewrite a pushed branch from
here. Force-pushing a rebased branch is only sanctioned under an explicit
user go-ahead, and even then the rebase alarm must come from the user, not
this skill.

## Step 0 — Preflight

```bash
git branch --show-current     # the child, when no branch argument is given
git remote get-url origin     # repo must have an origin remote
git symbolic-ref refs/remotes/origin/HEAD   # detect the default branch
```

- The **child** branch is the argument if one was given (a branch name, or
  "the branch we're on"), otherwise the current branch. An empty
  `git branch --show-current` means detached HEAD — **stop and ask** which
  branch to sync.
- No origin, or no remote HEAD → resolve the default branch by falling
  back to `main` (this repo's trunk is `main`), and **stop** if this repo
  isn't the one in play (wrong remote → don't guess).
- If child == the resolved default branch, **short-circuit**: the default
  branch syncs from its own remote, not from a parent. Report and stop.

## Step 1 — Identify the base branch

Resolve the parent in this order; take the first one that answers:

1. **Caller-provided base** — the user said "sync X with Y"; `Y` is the
   base. Trust it unless it's impossible (see the fork-point check below).
2. **Open-PR base** — if the child has an open PR, its base is the
   authoritative parent:
   ```bash
   gh pr view <child> --json baseRefName --jq .baseRefName
   ```
3. **Stored hint** — a base recorded at branch creation:
   ```bash
   git config --get branch.<child>.base || git config --get branch.<child>.mergeBase
   ```
   (Older conventions cache it there; present in this repo when a branch
   was created by tooling that records the fork point.)
4. **Default branch** — the repo's trunk, resolved in step 0
   (`origin/main` for everything cut straight off main, and for
   release-synced branches whose real parent is a `release/vX.Y` line,
   make the release branch explicit via path 1 or 2).

After picking a candidate, sanity-check the fork point:

```bash
git merge-base <child> <base>        # compare against the base's tip
git rev-parse <base>                 # if merge-base != tip, the base isn't the fork point
```

If the merge base isn't an ancestor relationship you'd expect for the
resolved base (i.e. the child's fork point is noticeably older/deeper than
the base tip), the child was likely cut from a *different* branch — a
feature parent or a release line. **Stop and ask** the user which parent to
sync against; don't guess a divergent lineage. `gh pr view` (path 2)
resolves exactly this when a PR exists.

## Step 2 — Fetch and survey the copies

```bash
git fetch origin --prune        # remote truth must be current before syncing
```

Then, for both the parent and the child, record what exists:

```bash
git rev-parse --verify --quiet refs/heads/<branch>   # local copy exists?
git ls-remote --heads origin <branch>                # remote copy exists?
```

## Step 3 — Reconcile the local parent with the remote

The goal is a single authoritative parent tip to update the child from.
Compute the delta:

```bash
git rev-list --left-right --count <parent>...origin/<parent>
# left = commits only in local parent; right = commits only in remote parent
```

| local parent        | remote parent        | action                                                                 |
| ------------------- | -------------------- | ---------------------------------------------------------------------- |
| missing             | exists               | `git branch --track <parent> origin/<parent>` (no checkout)             |
| exists, behind only | exists               | fast-forward local to remote; see below                                |
| exists, in sync     | exists               | nothing                                                                 |
| exists, ahead only  | exists               | nothing to *pull* — leave it; local parent is the newest tip           |
| exists, diverged    | exists               | **stop and ask** — local parent has unpushed commits *and* remote moved |
| exists              | missing              | warn (parent has no remote copy); keep using local                     |
| missing             | missing              | **stop** — base can't be reconciled; ask which parent to sync against   |

Fast-forward the behind-only case — `git branch -f` (`git merge --ff-only
origin/<parent>` if `<parent>` happens to be the checked-out branch,
because `git branch -f` refuses to move the current branch):

```bash
# behind-only, parent checked out:
git merge --ff-only origin/<parent>
# behind-only, parent not checked out:
git branch -f <parent> origin/<parent>
```

This is safe only because step 3 verified *behind-only* first — the local
is strictly behind, so the ref update is a pure forward. Verify after:
`git rev-parse <parent>` should equal `git rev-parse origin/<parent>`.

## Step 4 — Pick the authoritative parent tip

Now there is one newest parent tip to move the child onto:

- **Local parent exists and is ahead of (or equal to) remote** (unpushed
  local work, or remote missing) → use local `<parent>`. The child was
  cut from that lineage; syncing to the *older* remote would drop the
  unpushed commits from consideration.
- **Otherwise** → use `origin/<parent>` — the remote-tracking ref is the
  truth when local and remote agree, mirroring the branch skill's
  "anchor on `origin/<base>`, never the stale local" rule.

Call this `<src>` for the rest of the skill.

## Step 5 — Survey the child and move it onto `<src>`

Before touching the child, guard the working tree:

- **Dirty working tree** → **stop and ask**. Do not stash-and-continue
  unprompted; a mid-update stash is how work gets lost.
- Child != current branch → `git switch <child>` first (note the context
  switch in your report). If that fails, stop.
- Child == current branch (the common case) → proceed in place.

Then compute the child's relationship to `<src>`:

```bash
git rev-list --left-right --count <child>...<src>
# left = commits only on the child (ahead); right = commits the child is missing (behind)
```

Apply the type ladder from the top of this skill:

- **Behind 0, ahead > 0** — child is already ahead of its base. Nothing to
  pull; report "no parent updates available". Do not create a merge just
  to have one.
- **Behind > 0, ahead 0** — child is strictly behind → **fast-forward**:
  ```bash
  git merge --ff-only <src>
  ```
  Linear, no merge commit. (Forcing a `--no-ff` merge commit here is only
  worth it when the user explicitly wants a visible merge point.)
- **Behind > 0, ahead > 0** — diverged → decide merge vs rebase:
  - **Default: merge (merge commit)** — the preferred type, preserves both
    histories:
    ```bash
    git merge <src> --no-edit
    ```
    Only a real merge commit (ff is impossible once diverged).
  - **Rebase** — offer *only* when **both** hold: the child has **never
    been pushed** (no `origin/<child>` in the step-2 survey) **and** the
    user asked for/accepts a rebase. Then:
    ```bash
    git rebase <src>
    ```
    If the child **has** been pushed, rebasing would rewrite published
    history — the user must explicitly choose force-pushing it; the skill
    itself never proposes force, and only performs the rebase under that
    explicit go-ahead. `git rebase --onto <newbase> <oldbase> <child>` is
    the advanced variant for when the base lineage itself moved (rare;
    only when the user names two parents).
- **Equal** — already current. Report; don't fabricate an update.

**Conflicts** (merge or rebase): **stop**. Do not auto-resolve. Report the
conflict markers and state, and how to abort cleanly
(`git merge --abort` / `git rebase --abort`). The partial state is
deliberate — the user decides resolution.

## Step 6 — Verify and report

```bash
git status
git log --oneline --graph -5   # confirm the shape: ff, merge commit, or rebase
```

Confirm, before reporting success:

- The child now contains the parent tip: `git merge-base --is-ancestor
  <src> <child>` should succeed (for merge/ff) — or the user-confirmed
  rebase applied cleanly.
- Working tree is clean of the update (or shows only user-confirmed
  conflict leftovers).
- If the skill switched branches in step 5, say so, and state which branch
  is checked out now (offer to switch back rather than doing it silently).

Report: the identified base branch (+ how it was resolved), whether local
and remote copies existed for it, the update type applied, and the final
`<child>` vs `<src>` relation. Do not commit, push, or open a PR from here
— that is the commit-push-pr skill's job.

## Guardrails & judgment

- **Merge is the default, not the vestigial option.** Fast-forward to
  preserve linearity when it's free; a merge commit when history diverged
  — that's the whole ladder. Rebase is the exception with two hard gates
  (never-pushed child + explicit user confirmation).
- **Never rewrite pushed history on your own.** A pushed child that
  "needs a rebase" is a question for the user, never a course of action.
- **A dirty tree stops the show.** Don't ship someone's uncommitted work
  sideways.
- **The local parent must never be silently moved backward.** The only ref
  updates this skill performs on the parent are verified strictly-behind
  fast-forwards and fresh copies from remote — never a forced overwrite.
- **When the fork point disagrees, ask.** An unresolved lineage (feature
  off feature, or off a release line without a PR to confirm) is worth one
  question; guessing produces a wrong-looking diff.