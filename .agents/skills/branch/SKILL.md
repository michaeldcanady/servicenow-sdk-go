---
name: branch
description: >-
  Creates a new Git branch following modern branch-name practice and this repo's trunk-first
  flow: resolves the base (default branch, usually main), fetches so the base is fresh, derives
  a short kebab-case Conventional-Commit-style name oriented on the optional argument — an
  explicit branch name, a work description, or an issue reference (#123 / URL / title) — and
  verifies the switch. Guards: never deletes or recreates an existing branch (reuses it
  instead), never bases off a stale local branch, and flags a dirty working tree rather than
  silently moving changes. Use whenever the user asks to "create a branch", "start a branch
  for X", "make a new branch", or "get a branch for issue N"; also used internally by the
  commit-push-pr skill before committing.
---

# Creating a branch

A branch is the container for one unit of work, named so a future reader can
tell what it holds without opening it. This repo is trunk-first (ADR 011):
short-lived topic branches are cut from the trunk, merged via PR, and never
become a second development trunk. This skill creates one correctly — from a
fresh base, with a name that survives a `git branch` listing at arm's length.

## The optional argument

The argument is however the request was phrased — `$ARGUMENTS` on the command
line, an explicit branch name ("call it `fix/123-sysparm-view`"), an issue
reference (`#123`, a GitHub issue URL), or a plain work description ("branch
for adding the sysparm_view parameter"). Capture it in step 1; the rest of
the skill derives or validates the branch name from it.

## Step 1 — Capture the intent

Read the request/argument and classify it:

- **Explicit branch name** → use it verbatim after validation (step 4).
- **Issue reference** → resolve it first, so the name and base both reflect
  reality: `gh issue view <n> --json number,title,labels` (fall back to the
  bare number, title, or URL if `gh` isn't usable).
- **Work description** → this is the naming seed for step 3.

## Step 2 — Resolve the base branch and fetch it

The base is the repo's default branch unless the caller overrides it.
Resolve it the same way the rest of the repo does:

```bash
git symbolic-ref refs/remotes/origin/HEAD   # e.g. refs/remotes/origin/main
```

- Fall back to `main` if that fails (this repo's trunk is `main`). If the
  caller (e.g. the commit-push-pr skill) has already fetched and synced the
  base this session, reuse that state — don't re-fetch pointlessly.
- Otherwise make the base fresh before anchoring:
  ```bash
  git fetch origin --prune
  ```

Anchor the new branch on **`origin/<base>`**, the remote-tracking ref, not
the possibly-stale local `<base>`. A branch cut from a stale local base is a
merge-conflict PR waiting to happen.

## Step 3 — Derive the branch name

If the argument was an explicit name, skip straight to validation. Otherwise
derive one:

- **Conventional-Commit-style type prefix** (`<type>/`), matching the
  type vocabulary this repo's commits use (`feat`, `fix`, `docs`, `refactor`,
  `perf`, `chore`, `test`). Infer it from the request or issue — a bug/defect
  is `fix/`, a new capability is `feat/`, docs work is `docs/` — and only
  apply a prefix you can support from evidence. When the type is genuinely
  unknown, omit the prefix rather than guess.
- **A short kebab-case description**: lowercase, hyphen-separated, no spaces.
  From an issue use its title or `issue-<n>`; from a description, take the
  subject of the work, e.g. `add-sysparm-view-parameter`.
- **Issue number when available** — traceability beats cleverness:
  `fix/123-sysparm-view` or `docs/692-planning-hierarchy`.
- Full examples: `fix/123-sysparm-view`, `feat/add-attachment-upload`,
  `chore/ci/rebuild-docs-pipeline` (slash only as a separator, never in the
  description).

Keep it short and self-describing — the point is that `git branch` output is
readable, not that the name encodes a changelog.

## Step 4 — Validate the name

Reject any name that would break git or confuse tooling:

- Must not equal the default branch name or any existing tag.
- Git ref rules: no leading/trailing `.` or `-`, no `..`, no `~`, `^`, `:`,
  `?`, `*`, `[`, `\`, or spaces anywhere; no `@{`; no control characters.
  (At minimum, drop the description to lowercase kebab and if in doubt,
  prefer a shorter name.)
- Skip validation only for an explicit name the user dictated; still refuse
  anything the rules above reject — tell the user why rather than quietly
  renaming.

## Step 5 — Check for an existing branch first

Reuse, never recreate:

```bash
git branch --list <name>          # local
git ls-remote --heads origin <name>   # remote
```

- Branch exists locally → `git checkout <name>`, state that you reused it,
  and stop. Do **not** delete/recreate it — an existing branch may hold
  commits someone else relies on.
- Exists only on origin → checkout it and let git set up tracking
  (`git checkout <name>`, or `git branch --track <name> origin/<name>`).
- Otherwise the branch is new — proceed to create it.

## Step 6 — Create the branch off the fresh base

```bash
git switch -c <name> origin/<base>
```

- If the working tree is dirty: `git switch -c` will carry uncommitted
  changes over to the new branch, which is *usually* the intent when someone
  says "branch for this work" — but state clearly to the user that the
  changes moved with the branch. If the request was purely administrative
  ("create a branch, nothing else") and the tree is dirty, **stop and ask**
  rather than silently relocating their work. Never stash-and-drop.
- Verify the switch landed: `git branch --show-current` must print the new
  branch name, and `git status` should confirm the working tree is as
  expected (clean, or carrying the changes if that was the intent).

## Step 7 — Report

Report the branch name, the base it was cut from (and its commit, e.g.
`origin/main` at `c97e1126`), and whether the working tree carried changes
along. Do not commit, push, or open a PR from here — those are the
commit-push-pr skill's steps.