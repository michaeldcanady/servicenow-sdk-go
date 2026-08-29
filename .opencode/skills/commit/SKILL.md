---
name: commit
description: >-
  Writes well-formed git commits following modern best practices and this repo's
  release-please / Conventional-Commits conventions. Inspects the working tree first, stages
  only the intended files, splits work into atomic logical commits, and writes an imperative
  subject under 50 characters with a why-focused body wrapped at 72. Picks the type and any
  BREAKING CHANGE footer from the actual diff (feat/fix/docs/refactor/perf/chore/test; CI and
  workflow changes are always chore). Adds the standard Co-authored-by trailer per repo
  convention. Never stages secrets or unrelated files, never skips hooks or force-pushes, and
  never pushes. Use whenever the user asks to "commit", "stage and commit", "make a commit",
  "write a good commit message", or after a unit of work is finished and a commit is requested.
---

# Writing git commits

A commit message is read more times than it is written: by reviewers, by
`git bisect`, by the person who breaks the build next quarter, and — in this
repo — by release-please, which parses every message to decide what the next
version and `CHANGELOG.md` entry look like. A good message lets someone
decide "do I need to look at this diff?" without looking at the diff. This
skill is the checklist for producing one.

These rules combine the modern baseline for commit quality (the Conventional
Commits spec and the classic "50/72 + imperative + explain why" rules) with
this repo's specific versioning contract. The repo rules win when they
stray from the general guidance — they exist because release automation
depends on them.

## When to commit

Only commit when the user or coordinating session explicitly asked for a
commit. "Done with the work" is not a commit request. Never push unless
asked, and never open a PR unless asked.

## Step 1 — Inspect before staging

```
git status
git diff            # unstaged changes
git diff --cached   # already-staged changes
git log --oneline -10   # match the repo's message style
```

Goal: know exactly what changed, what is staged, and what the established
message shape looks like. Never commit from memory of what you intended to
change — confirm it against the working tree.

## Step 2 — Make it atomic

A commit is a unit of **meaning**, not a unit of time. Split or group work so
each commit is one logical change: one bug, one feature, one refactor, one
fix. Rules of thumb:

- If a change spans multiple unrelated concerns (e.g. a fix plus a drive-by
  rename), split them — stage different files (or different hunks with
  `git add -p`) into separate commits.
- Stage only the files you actually touched. Never `git add -A` or
  `git add .` on faith, and never include files outside the task's scope.
- Never commit secrets, credentials, `.env` files, or build artifacts —
  review `git diff --cached` before committing to confirm.
- Tiny atomic commits ("fix typo", "rename variable") are fine and normal.
  The smell is a commit that bundles many concerns, or a vague catch-all.

If existing staged changes don't match a clean unit of work, unstage
(`git restore --staged`) what doesn't belong and commit only the coherent
part — don't force one message to cover everything.

## Step 3 — Choose the type from the diff

Subject shape: `<type>(<scope>): <subject>`

Types used in this repo (each drives release-please's category rules, so
pick by what the diff actually does, not by convenience):

| Type | When |
| --- | --- |
| `feat` | A new user-visible capability (new endpoint, new module, new option) |
| `fix` | A bug fix — behavior changes from wrong to right |
| `docs` | Documentation only: `website/docs/`, package `Readme.md`, comments |
| `refactor` | Behavior-preserving code change (renames, restructuring) |
| `perf` | Performance improvement with no behavior change |
| `chore` | Maintenance: dependencies, config, scaffolding |
| `test` | Adding/fixing tests or test infrastructure |

Repo-specific rules that override general guidance:

- **CI/workflow changes are always `chore:`** — never `fix:`/`feat:`, even
  when they fix a broken run. Anything touching `.github/workflows/`,
  CI-only `scripts/`, or pipeline plumbing must not appear in `CHANGELOG.md`
  as a fix or feature.
- **Never hand-edit `VERSION` or `CHANGELOG.md`** — release-please generates
  both from commit messages.
- Scope (optional): use the package/area name when it helps narrow the
  subject, e.g. `feat(tableapi): ...`, `docs(website): ...`, `chore(ci): ...`.

## Step 4 — Write the message

### Subject (first line)

- Imperative mood, present tense: "Fix", "Add", "Refactor", not "Fixed",
  "Adds", "Refactoring".
- Keep it under 50 characters — if it wraps, tighten it. Don't end with a
  period.
- Name the component and the change: `fix(tableapi): return error on empty
  sys_id` not `fix: stuff`.
- Under Conventional Commits, write the description in lowercase after the
  type or scope — all-lowercase is the spec's convention and matches the
  majority of this repo's history. The sentence-case rule (capitalize the
  first word) is for plain, non-Conventional Commit subjects.

### Body

Separate from the subject with a blank line. Wrap at 72 characters.

The body explains **why**, not **what** — the diff already shows what.
Write it like a journalist writing a lead: state what happened and why it
mattered, in as few words as possible. Assume the reader (your future self
included) has none of your context — the change that's obvious to you now
will not be obvious in six months, so "tell, don't [just] show". No filler
words ("though", "maybe", "I think", "kind of").

Include in the body, as relevant:

- The problem being solved and why it mattered.
- The trade-offs or alternatives considered.
- Behavior contrast: what changed from before, and any compatibility
  impact (a published SDK's breaking changes are a semver event).
- Breaking changes: a `BREAKING CHANGE: <description>` footer (or `!` after
  the type/scope) explaining old vs. new behavior — this is what drives a
  major version bump in this repo's release automation.
- References: `Fixes #123`, `Closes #456`, `Refs #789` as a footer.

A short subject-only commit is fine when a one-liner fully captures the
change — never pad the body just to have one. But when the change has
context a reader can't recover from the diff, put it in the body.

### Example

```
feat(tableapi): add sysparm_view query parameter

Exposes ServiceNow's sysparm_view for table queries so callers can get
alternate view formatting without an extra round trip. Mirrors the existing
sysparm_display_value wiring; the parameter is optional and defaults to
unset as before.

Closes #123
Co-authored-by: opencode <opencode@local>
```

## Step 5 — Commit

Stage the intended files explicitly, then verify exactly what is staged once
more (`git diff --cached`) before committing:

```bash
git add <path/to/file.go> <path/to/test.go>
git commit -m "fix(tableapi): return error on empty sys_id" \
           -m "Explain why here, wrapped at 72 columns."
```

Use one `-m` per paragraph; for long messages prefer `-m "$(cat <<'EOF' ... EOF)"`
or a `--file` temp file over fragile inline newlines.

Repo commit conventions on top of the mechanics:

- Include the `Co-authored-by: opencode <opencode@local>` trailer on
  commits you author, matching the existing history.
- Do not skip hooks (`--no-verify`), do not use interactive rebase
  (`-i`) to rewrite, do not force-push, do not create empty commits —
  unless the user explicitly asks.
- If a commit or its hooks reject the change, fix the problem and make a
  **new** commit — do not `--amend` the failed commit.
- Do not amend or rewrite a commit that is already pushed.

## Step 6 — Verify

```bash
git log -1 --stat    # or: git show --stat HEAD
git status           # should be clean of the just-committed files
```

Confirm the message reads correctly, the type produced the right
release-category behavior, and nothing unintended slipped in. Report the
commit hash and subject to the user. Pushing, PRing, or tagging are separate
explicit asks.

## Anti-patterns

- "Update stuff", "wip", "changes" — vague subjects that force a diff read.
- Logging the changes ("Changes: fixed X, changed Y...") — that's what the
  diff is for; the body should say why.
- Hedged, filler-laden subjects ("I think I fixed it", "maybe fix query") —
  the imperative tells the future reader what happened; hedges just burn
  the subject line.
- Bundling unrelated fixes into one commit so a release note reads as one
  feature.
- Padding a header on every commit regardless of content ("Generated by
  tool X" boilerplate adds no signal).
- Committing files you didn't intend to (secrets, `coverage.html`,
  generated build output) — caught at Step 1/5 review.

## Further reading

- Conventional Commits spec: https://www.conventionalcommits.org/en/v1.0.0/
- freeCodeCamp — How to Write Better Git Commit Messages:
  https://www.freecodecamp.org/news/how-to-write-better-git-commit-messages/