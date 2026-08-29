---
name: commit-push-pr
description: >-
  Takes the current uncommitted working-tree changes all the way to a GitHub PR: stashes them,
  syncs the local default branch (main) with origin, restores the changes on a PR-capable
  branch, commits them via the commit skill, pushes the branch, and opens the PR against
  the default branch. Accepts an optional issue reference; when no issue is provided the PR is
  labeled no-issue-required. Use whenever the user asks to "commit, push, and open a PR", "ship
  these changes", "commit and PR this", or similar — never use on a clean tree with nothing to
  ship, and never push to the default branch directly.
---

# Commit → Push → PR

This skill runs the working-tree change set through the full ship path:
stash → sync trunk → restore on a feature branch → commit → push → open a
PR. It is the outer loop that hands the actual branch creation to the
**branch** skill (step 5) and the commit-writing to the **commit**
skill (step 6) — both own their conventions and safety rails, so don't
duplicate their rules here.

The gating rule before anything else: **never ship nothing**. Step 1
short-circuits a clean tree, and step 8 refuses to push to the default
branch.

## The optional issue argument

This skill accepts an issue reference as an argument. It arrives however the
request was phrased — `$ARGUMENTS` on the command line, an issue number like
`#123`, a GitHub issue URL, or a referenced issue title. Parse the number (or
URL) out of the request.

- **Issue provided** → link the PR to it: `Closes #<n>` in the PR body, and
  name the branch from it (see step 5).
- **No issue provided** → the PR **must** be labeled `no-issue-required`
  (step 8). This is the explicit fallback the repo wants when work ships
  without an issue on file.

## Step 0 — Preflight

Before touching anything, confirm the pipeline exists:

```bash
git remote get-url origin      # repo must have an origin remote
git symbolic-ref refs/remotes/origin/HEAD   # detect default branch
gh auth status                 # gh must be authenticated for the PR step
```

- Resolve the default (trunk) branch: `git symbolic-ref
  refs/remotes/origin/HEAD` gives e.g. `refs/remotes/origin/main`
  (fall back to `main` if the command fails — this repo's trunk is `main`).
- If any preflight fails (no origin, no remote HEAD, gh not authed), **stop**
  and tell the user what's missing. Don't guess the default branch.

## Step 1 — Check for uncommitted changes

```bash
git status --porcelain
```

- Non-empty → there is something to ship; continue to step 2.
- Empty → **short-circuit**: there is nothing to commit, push, or PR. Report
  "no uncommitted changes — nothing to ship" and stop. Do not invent work.
  (If the tree is clean but the current branch has *unpushed commits*, mention
  that and offer to push + PR them as an explicit opt-in — don't proceed
  automatically.)

## Step 2 — Stash the changes

```bash
git branch --show-current     # record the branch you're on for step 5
git stash push -u -m "commit-push-pr: <one-line summary of the work>"
```

- Use `-u` so untracked new files ride along with the stash — a plain stash
  would leave them loose and they could collide when the tree is restored.
- Guard the stash until the end: **never `git stash drop` before the PR is
  confirmed open** (step 9), so a mid-flow failure never destroys the work.
- After stashing, the working tree is clean, which is what makes steps 3–5
  safe.

## Step 3 — Change to the default branch

```bash
git fetch origin --prune
git checkout <default>                # creates local tracking branch if missing
```

If local `<default>` doesn't exist yet, `git checkout -b <default>
origin/<default>` to create it tracking the remote. Prefer `git switch`/`git
checkout` over anything that resets branch state.

## Step 4 — Make sure local is up to date with remote

```bash
git merge --ff-only origin/<default>   # or: git pull --ff-only
```

Fast-forward only. If the merge refuses (local `<default>` has commits the
remote doesn't), **stop and ask** — the local default branch has diverged and
rewriting or force-flagging it is not this skill's call to make. Never create
a merge commit just to reconcile the trunk.

## Step 5 — Restore the stashed changes on a PR branch

You can't open a PR with head == base; the changes cannot be committed on the
default branch itself. Decide where they land:

- **Reusable branch**: if the pre-stash branch (from step 2) is a real,
  non-default branch, check it out again and re-base it onto the fresh trunk
  so the PR contains no stale history:
  ```bash
  git checkout <pre-stash-branch>
  git rebase origin/<default>
  ```
  **But only if that branch hasn't been pushed.** Check first:
  ```bash
  git ls-remote --heads origin <pre-stash-branch> | grep -q .
  ```
  Rebasing a pushed branch rewrites every commit hash on it, which step 7's
  never-force-push rule can't ship — a normal push would be rejected as
  non-fast-forward and force is forbidden, dead-ending the flow. If the
  branch exists on origin, **stop and ask** the user to choose: rebase with a
  one-time `--force-with-lease` push under their explicit go-ahead (the only
  sanctioned exception to the never-force rule), skip the rebase and keep the
  branch as-is, or fall through to the fresh-branch path below. Note: the
  fresh-branch path carries only the stashed (uncommitted) work — any commits
  the user previously pushed on the old branch stay behind and are excluded
  from the new PR, so that choice orphaning real work. Never rewrite
  a pushed branch's history without that decision.
- **Fresh branch**: if you were on the default branch, detached HEAD, or no
  branch at all, run the **branch** skill to cut a branch off the
  freshly synced trunk. Hand it the base (`<default>`) and the naming seed —
  the issue reference (`#123`) when there is one, otherwise a short work
  description. It owns name derivation, git-ref validation, and the
  existing-branch check, so don't duplicate that logic here.

Then restore the work, **on that branch**:

```bash
git stash pop
```

- If `git stash pop` conflicts, **stop**. Do not auto-resolve, do not drop
  the stash. Surface the conflict markers and the current state to the user.
  The stash stays intact until they decide.

## Step 6 — Commit via the commit skill

Run the **commit** skill and follow its process end-to-end (inspect →
atomic → type from the diff → write subject/body → stage explicitly →
commit → verify). That skill owns subject length, type selection (including
the always-`chore:` CI rule), the `Co-Authored-By` trailer, and the
no-force/no-hook-skip rails.

If an issue was provided, reference it in the commit body (`Closes #123`) —
the release process and the PR both benefit from the link living at the
commit.

## Step 7 — Push the changes

```bash
git push -u origin <branch>
```

- The command above is a hard rule: **never force-push** (`--force`,
  `--force-with-lease`), and never push to the default branch.
- If the push is rejected as non-fast-forward (someone updated the branch),
  `git fetch origin` then `git rebase origin/<default>` (or `origin/<branch>`
  if that's what moved) and retry a normal push — never a forced one.

## Step 8 — Open the pull request

```bash
gh pr create --base <default> --head <branch> \
  --title "<commit subject from git log -1 --pretty=%s>" \
  --body "<commit body from git log -1 --pretty=%b>"
```

- **Issue provided**: append the `Closes #<n>` link to the body.
- **No issue provided**: the PR **must** carry the `no-issue-required` label:
  ```bash
  gh pr create ... --label no-issue-required
  ```
  If that fails because the label doesn't exist on the repo, create the PR
  first and then `gh pr edit <number> --add-label no-issue-required`; if that
  also reports the label is missing, stop and tell the user the label may
  need to be created — a missing label is a repo-admin fact, not something to
  silently paper over.

Let `gh` also fill in the default body template when the commit body is
empty — an empty-handed `--body ""` forfeits the repo's template fields.

## Step 9 — Verify and clean up

```bash
gh pr view <number> --json url,title,baseRefName,headRefName,labels
git status
git stash list
```

Confirm, before reporting success:

- The PR exists with the right base/head and correct title.
- The `no-issue-required` label is attached **when no issue was given**.
- The working tree is clean and the branch is pushed (`git status` should
  show "up to date with origin/<branch>").
- No stash remains: if `git stash list` still shows one, the pop didn't fully
  restore — report it and **do not drop it** without the user's go-ahead.

Report the PR URL and a one-line summary. You may leave the branch checked
out — switching back to the default branch is not automatic.

## Abort rules

Stop and surface state — never work around it silently:

- Step 1: clean tree → nothing to ship.
- Step 4: local default branch diverged from origin → ask, don't rewrite.
- Step 5: `git stash pop` conflicts → stop, keep the stash, show the
  conflicts.
- Step 7: push rejected → rebase then retry; never force.
- Step 8: missing `no-issue-required` label → report; don't fake the label.
- Any failure before the PR is open → the work is still safe in the stash;
  say so explicitly in your report.