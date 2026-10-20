---
name: write-issue
description: >
  Files a well-formed GitHub issue in the current repository, using whatever issue templates the
  repo actually defines instead of a generic bug/feature guess. Use this whenever the user wants to
  "open an issue", "file a bug", "report this as an issue", "write up a feature request", "create a
  GitHub issue for X", or after a bug/design discussion has concluded and the user wants it tracked
  formally rather than just fixed in place. Also use it proactively when the user describes a defect
  or missing capability and asks you to "track this" or "make sure this doesn't get lost" — that's
  an issue-filing request even without the word "issue" in it. Do not use this for opening pull
  requests, editing existing issues, or commenting on issues — only for creating new ones.
---

# write-issue

## Why this exists

Every repo's issue tracker has its own shape: some have no templates at all, some have GitHub Issue
Forms (YAML, with dropdowns and required fields), some have old-style markdown templates, and almost
all of them care about labels being applied correctly so triage isn't manual work. Guessing at this —
or always producing the same generic "## Description / ## Steps to reproduce" body regardless of what
the repo actually asks for — produces issues that don't match the project's conventions and often
miss required fields a maintainer will just have to chase down. This skill's job is to read the repo's
actual conventions before writing anything, the same way a contributor who'd done their homework would.

## Workflow

Work through these steps in order. Don't skip the discovery steps even if you're confident you
remember a repo's conventions from earlier in the conversation — repos change, and getting the repo
name or template shape wrong produces an issue in the wrong place or missing required fields.

### 1. Resolve the repository

Run `git remote get-url origin` (or the appropriate remote if `origin` isn't a GitHub URL — check
`git remote -v` if in doubt). Parse `owner/repo` out of either the SSH (`git@github.com:owner/repo.git`)
or HTTPS (`https://github.com/owner/repo.git` or `https://github.com/owner/repo`) form.

If there's no git repo, no remote, or the remote isn't a `github.com` URL, ask the user for the
`owner/repo` directly rather than guessing — don't file an issue against the wrong repository.

### 2. Confirm issues are enabled

Before doing anything else with the repo, check whether Issues is actually turned on:

```bash
gh repo view <owner>/<repo> --json hasIssuesEnabled,isArchived
```

- If `hasIssuesEnabled` is `false`, or the repo is archived, **stop here**. Tell the user plainly
  that Issues isn't enabled (or the repo is archived) and that you can't file anything — don't fall
  back to opening a discussion, a PR, or anything else instead unless the user explicitly asks you to.
- If the `gh` CLI isn't authenticated or the repo isn't reachable, say so and stop rather than guessing
  at repo state.

### 3. Discover issue templates

Check, in order:

1. `.github/ISSUE_TEMPLATE/` — the standard location. List its contents. Each `.yml`/`.yaml` file is
   a GitHub Issue Form; each `.md` file is a legacy markdown template. There's often also a
   `config.yml` alongside them — read it, since it can disable the blank/no-template option
   (`blank_issues_enabled: false`) or point to `contact_links` (e.g. "search existing issues first",
   "use Discussions for questions") that change whether filing an issue is even the right move.
2. A single `.github/ISSUE_TEMPLATE.md` — older single-template convention.
3. Neither exists — the repo has no templates. That's not an error condition, just fewer constraints.

Read each template file directly (they're short) rather than trying to parse them mechanically —
YAML issue forms vary in structure and are easier to interpret by reading than by pattern-matching.
For each template, note: its `name`, `description`, `title` prefix, `labels`, and — for YAML forms —
the list of fields (`label`/`id`, whether `required`, and any `dropdown`/`checkboxes` options), since
you'll need to answer each required field to produce a well-formed issue body.

Some repos (this one included, as of when this skill was written) keep a second tier of templates for
maintainer-only planning artifacts (e.g. Epics/Tasks) outside `ISSUE_TEMPLATE/`, referenced from a
`config.yml` comment or `CONTRIBUTING.md` rather than exposed in the GitHub "New Issue" chooser. If
you spot something like that and the user's request matches it better than the public templates
(e.g. they're clearly scoping multi-story work, not filing a single bug), mention it and ask which
they want rather than silently picking the public template.

### 4. Decide what kind of issue this is

Match the substance of what the user described against the templates you found, not against generic
categories. If the repo's templates are named "Bug Report" and "Feature Request", a broken-behavior
report is a bug and a "we should support X" is a feature — but don't force a fit: if the user's
request doesn't cleanly match any template (e.g. it's a question, a chore, or something templates
don't cover), say so and ask them how they'd like to proceed rather than jamming it into the closest
template and leaving required fields blank or nonsensical.

If there are no templates at all, this step is simple: write a clear title and a body with whatever
sections make sense for the content (what's happening, what's expected, repro steps if it's a bug,
etc.) — you don't need to invent a template structure the repo doesn't have.

### 5. Check for duplicates

Before drafting anything the user has to review, search existing issues (open and closed) for likely
duplicates:

```bash
gh issue list --repo <owner>/<repo> --search "<keywords>" --state all --limit 10
```

(or `mcp__plugin_github_github__search_issues` if the GitHub MCP server is available — prefer it when
present, since it integrates with the rest of this conversation's tool use more smoothly). Use a few
different keyword combinations pulled from the user's description, not just one literal phrase — the
same bug is often filed with very different wording.

If you find something that looks like a real duplicate (not just a loosely related issue), tell the
user what you found and ask whether they still want to file a new one, want you to comment on the
existing issue instead, or want to drop it. Don't create the new issue without asking in this case —
a duplicate that goes in anyway makes triage strictly harder, not easier. If nothing close turns up,
just mention briefly that you checked and move on — no need to make the user confirm a negative.

### 6. Draft the issue

Fill in the template's required fields using the conversation context you already have. For each
field the conversation doesn't already answer, ask the user rather than inventing plausible-sounding
content — a fabricated "Steps to reproduce" or version number is worse than an honest gap, since it
actively misleads whoever triages it.

For YAML issue forms, reproduce the form's structure as markdown headers matching each field's
`label`, in the same order as the form — this is what GitHub itself renders when a form is submitted
through the UI, so the resulting issue looks native rather than hand-rolled. Apply the template's
`labels` and use its `title` prefix.

Keep the title itself short — name the symbol/component and the defect, not the full explanation.
The "why it matters," root cause, and suggested fix belong in the body (the template already has
fields for them); a title that tries to carry that too becomes a wall of text in issue lists, search
results, and commit references. `[Bug]: (*NilPointerError).Error() panics on a nil receiver` is a
title; `[Bug]: (*NilPointerError).Error() panics on a nil receiver instead of returning a safe
string` is a title plus a sentence of the "what did you expect" field bolted on.

Show the user the drafted title, labels, and body before creating anything, and get their go-ahead —
filing an issue is a visible, hard-to-cleanly-undo action (closing it isn't the same as it never
having existed), so treat it the way you'd treat any other action with an external, shared effect.
Skip this confirmation only if the user has already explicitly told you to just go ahead and create it.

### 7. Create the issue

Prefer the GitHub MCP server's issue-creation tool if it's available in this session (check for
`mcp__*github*__issue_write` or similarly named tools) — it validates labels against the repo's
actual label set and gives structured errors instead of a CLI exit code. Fall back to `gh`:

```bash
gh issue create --repo <owner>/<repo> --title "<title>" --body "<body>" --label "<label1>" --label "<label2>"
```

Use `--body-file` with a temp file instead of `--body` if the body is long or contains characters
that are awkward to escape inline.

### 8. Report back

Give the user the issue URL and a one-line summary of what was filed (title + labels). Don't
re-paste the full body back to them — they just reviewed it in step 6.

## Notes on judgment calls

- If the user's request is really several distinct issues (e.g. "these three things are all broken"),
  ask whether they want one issue or several before drafting — don't default to bundling or splitting
  without checking, since either can be wrong depending on how the maintainer triages.
- If a required template field asks for something you can determine yourself from the repo (e.g. an
  "SDK Version" field you can get from `go.mod` or a package manifest, a Go/Node/etc. version from the
  environment), fill it in rather than asking — only ask for things that genuinely require the user's
  input.
- If issues are enabled but the repo's `config.yml` sets `blank_issues_enabled: false` and nothing
  matches the user's request, tell them the templates don't cover this and ask how they'd like to
  proceed (closest-fit template with a note, or file through the web UI themselves) rather than
  silently ignoring the constraint.
