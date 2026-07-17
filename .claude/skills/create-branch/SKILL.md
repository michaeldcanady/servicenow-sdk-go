---
name: create-branch
description: Create a correctly-named feature branch in this repo before committing new work — picks a type/kebab-description name matching this repo's convention (e.g. fix/tableapi-nil-pointer, docs/error-sentinel-notes), and checks whether a new branch is even needed. Use whenever the user asks to "create a branch", "start a new branch", "branch off main", or whenever another workflow (like opening a PR) needs a feature branch before committing directly to main or another shared branch.
---

# create-branch

This repo's branches follow `<type>/<short-kebab-description>` — the same
`type` vocabulary as Conventional Commits (`feat`, `fix`, `refactor`, `docs`,
`perf`, `chore`, `test`), not a generic name or a ticket number. Recent
examples from this repo's history: `fix/paging-link-headers`,
`feat/account-api`, `refactor/model-accessor-mutators-use-store`,
`chore/ci-adjusting-integration-testings`. A branch named `my-feature` or
`update-stuff` doesn't fit that pattern and makes the eventual PR title
(which the [[pr-workflow]] skill also derives from Conventional Commits)
harder to land on consistently.

## Step 1 — check whether a new branch is actually needed

Run `git branch --show-current` and `git status`. Don't create a branch
reflexively:

- If the current branch is already a feature branch (not `main`, not
  `release/*`, not some other shared/trunk branch) and it makes sense to
  carry on there, there's nothing to do — say so and stop.
- If the current branch is `main`, `release/*`, or another shared branch,
  a new branch is required before anything gets committed to it — shared
  branches should never receive direct commits.
- If unsure whether an existing branch is "shared," check `git branch -a`
  for remote tracking of it, or ask.

## Step 2 — pick the type

Read the diff (`git diff`, plus `git status` for untracked files) or, if
nothing is written yet, ask what the change will do. Pick `<type>` the same
way the [[conventional-commit]] skill picks a commit type — by what the
change *does*, not by how the user phrased the request:

| Type | Use when |
|---|---|
| `feat` | New capability being added |
| `fix` | A bug is being corrected |
| `refactor` | Code reshaped, no behavior change |
| `docs` | Only docs/comments/Readme changes |
| `perf` | Change specifically for speed/memory |
| `chore` | Tooling, CI, dependency, config work |
| `test` | Only test files, no production code |

If the work doesn't exist yet (branching ahead of writing any code), pick the
type that matches the user's stated intent — you can't read a diff that
doesn't exist yet.

## Step 3 — pick the description

A few kebab-case words describing the change, not a restatement of the type
and not a ticket number unless the user gives one. Check `git branch -a` if
you want to match this repo's existing naming texture, and make sure the name
isn't already taken locally or on `origin`.

## Step 4 — create it

```
git checkout -b <type>/<short-kebab-description>
```

Branch off the correct base — usually `main`, but if the user is working
against a release branch (e.g. `release/2.0`) or another integration branch,
branch from that instead. `git branch --show-current` before this step tells
you what you're branching from.

Report the branch name back so the user (or the calling workflow) knows what
was created.
