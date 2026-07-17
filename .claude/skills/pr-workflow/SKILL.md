---
name: pr-workflow
description: Take a change from working tree to an open GitHub pull request in this repo — branch, commit (Conventional Commits), push, and gh pr create with a lint-passing title/body — then check CI status on request. Use whenever the user asks to "open a PR", "make a pull request", "push this up", "get this ready for review", or asks about the status/checks of a PR they opened, even if they only ask for one step of the sequence (e.g. just "commit and push this"). Always confirm the PR title/body and never merges automatically — merging stays a manual, explicit action.
---

# pr-workflow

Carries a change through this repo's actual PR path: branch → commit → push →
`gh pr create` → (later) check status. Nothing here merges a PR — that step is
always left to the user, since it's a hard-to-reverse, shared-state action.

## Why this shape, not a generic PR skill

This repo enforces two things a generic "open a PR" flow would miss:

- **`.github/workflows/pr.yml`** runs `amannn/action-semantic-pull-request` on
  every PR — the **title** must itself be a valid Conventional Commit header
  (`type(scope): description`, scope optional). A title like "Fix bug" fails
  CI before a human even looks at it.
- **Branch names in this repo's history** follow `type/short-kebab-description`
  (e.g. `fix/paging-link-headers`, `feat/account-api`,
  `refactor/model-accessor-mutators-use-store`) — mirroring the commit type,
  not a generic `my-feature` or a ticket number.

## Step 1 — figure out where things stand

Run `git status`, `git branch --show-current`, and `git log main..HEAD --oneline`
(or `git log <base>..HEAD` if the base isn't `main`) to see:
- whether there's uncommitted work
- whether the current branch is `main`/`release/*` (shared branches — never
  commit directly to these) or already a feature branch
- whether there are already commits ahead of the base that just need pushing

Don't assume — a user asking to "open a PR" might already be three commits
into a feature branch with nothing to commit, or might be sitting on
uncommitted changes on `main`.

## Step 2 — branch (only if needed)

If the current branch is `main`, `release/*`, or otherwise not meant to carry
this change, create a new branch off the base before committing anything:

```
git checkout -b <type>/<short-kebab-description>
```

Pick `<type>` the same way the [[conventional-commit]] skill picks a commit
type (`feat`, `fix`, `refactor`, `docs`, `perf`, `chore`, `test`) — read the
diff, don't just echo the user's phrasing. `<short-kebab-description>` should
be a few words, matching the style of existing branches (check `git branch -a`
if unsure).

If a suitable feature branch already exists and is checked out, skip this.

## Step 3 — commit

Invoke the `conventional-commit` skill for this — it already encodes this
repo's type/scope/description rules and the git-safety norms (stage files by
name, heredoc message, no `--no-verify`, no amend unless asked). Don't
re-derive commit-message rules here; that skill is the source of truth for
them.

If there's nothing uncommitted (Step 1 already showed pushable commits),
skip straight to Step 4.

## Step 4 — push

```
git push -u origin <branch>
```

(omit `-u` if the branch already has an upstream). This is a shared-visibility
action but a routine and easily-reversible one for a feature branch the user
is actively working on — proceed without a separate confirmation prompt
unless something looks unusual (e.g. pushing to someone else's branch).

## Step 5 — draft title and body, then confirm before opening

**Never call `gh pr create` without showing the title and body first and
getting explicit go-ahead.** A wrong scope or a vague summary is cheap to fix
before the PR exists and annoying to fix after (title edits re-trigger the
lint check and notify watchers).

**Title**: must satisfy the semantic-pull-request check — `type(scope): desc`
or `type: desc`, same `type` list as commits. If the branch carries one
commit, the title is usually that commit's message verbatim (it already went
through the same rules). If it carries several commits, don't just concatenate
them — read the actual diff (`git diff main...HEAD`) and write one title that
describes the net effect, the same way you'd pick a single type for a mixed
commit in `conventional-commit`.

**Body**: this repo has no `.github/pull_request_template.md`, so build a
short one from the diff and commit log rather than leaving it blank:

```markdown
## Summary
- <1-3 bullets on what changed and why, from the diff/commits — not a
  restatement of the title>

## Test plan
- <what was actually run: `go test ./...`, a specific package's tests,
  manual verification>
```

Before writing the Test plan bullets, actually run something — don't guess
what a reviewer would want to see. If the change touches `.go` files, run
`go build ./...` and `go test ./<touched-package>/...` (or the full suite for
a small change) right now if you haven't already this session, and report
the real result ("passes", or "pre-existing failure in X, unrelated to this
change — confirmed via `git stash`" if something's already broken on the
base branch). For a docs-only or config-only change, it's fine to say so
explicitly instead of running Go tests. A Test plan that just says "no tests
run" or "manual review" without having tried the obvious command is a worse
signal to a reviewer than not having a PR at all — it looks like the change
was never checked.

Keep it as short as the change warrants — a one-line dependency bump doesn't
need three bullets under Summary.

Show the drafted title + body to the user and wait for approval or edits
before proceeding to Step 6.

## Step 6 — open the PR

```
gh pr create --base main --title "<approved title>" --body "$(cat <<'EOF'
<approved body>
EOF
)"
```

Use `--base main` unless the user is targeting a different base (e.g. a
release branch) — check `git log --oneline -5` against remote branches if the
base isn't obvious. Report the returned PR URL back to the user; don't open it
in a browser unless asked.

If a `gh pr create` prompt asks about forking/permissions and this isn't the
user's repo, stop and ask rather than guessing.

## Checking status (separate ask, not always part of the flow)

When the user asks about a PR's status, checks, or whether it's ready —
whether or not this skill just opened it:

```
gh pr checks <number-or-branch>
gh pr view <number-or-branch> --json mergeable,mergeStateStatus,statusCheckRollup
```

Report what's passing/failing/pending in plain terms (e.g. "lint-title and
unit-tests are green, CodeQL is still running"). If everything is green and
mergeable, say so — but **do not run `gh pr merge`**. That's the one step in
this whole flow that stays in the user's hands; if they want it merged, they
say so explicitly and you confirm the merge method (squash/merge/rebase) they
want, matching whatever this repo's branch protection expects.
