---
name: check-related-issue
description: >-
  Checks whether an already-open GitHub issue relates to a unit of work about to be shipped: it
  searches the repo's open issues with keywords pulled from a branch name, commit subject, stash
  message, or work description, scores the candidates, and declares a verdict (a strong related
  issue, ambiguous candidates, or none — plus a flag when the work is already in flight on an
  open PR). Use before opening a PR (the commit-push-pr skill runs it automatically), before
  starting a branch for work, or whenever the user asks "is there an issue that already tracks
  this?", "check for a related open issue", or "does an issue exist for X". Do not use it to file
  issues (that is write-issue) or to dedupe newly-filed report issues (that is the CI
  issue-similarity-check workflow — this skill looks backwards from work about to ship, not
  forwards from a fresh report).
---

# check-related-issue

## Why this exists

A PR carries one of two signals about issue tracking: a link to an issue
it closes (`Closes #n`), or the `no-issue-required` label. The second
signal is only honest if no open issue actually tracks the work —
otherwise the ship path silently duplicates a tracked item, the issue is
never closed by the PR, and `branch-policy`'s "issue or
no-issue-required" rule is satisfied by accident while triage loses the
link. This skill closes that gap: given the work about to ship, it finds
out whether an open issue already describes it, so the caller can link it
instead of defaulting to `no-issue-required`. It is deliberately
conservative — a false positive (linking an unrelated issue)
is worse than a miss, because linking claims the PR resolves that issue.

## When to use

- Automatically, from the **commit-push-pr** skill, just before the PR is
  opened (it hands this skill the work description and reads the verdict).
- Directly, when the user wants to know whether an issue already covers
  work they're about to branch or PR.
- As a lighter confirmation when the user did provide an issue reference
  but you want to make sure the work isn't *also* tracked elsewhere and
  isn't already in flight on another PR.

## Inputs

You need something that describes the work. The caller passes as much of
this as it has, best first:

- The user's request phrasing (what they asked to ship).
- The branch name (a `fix/123-add-sysparm-view` branch encodes both an
  issue number and a description).
- The commit subject of the last commit on the branch (most concrete when
  it exists).
- The stash/work-summary message (the commit-push-pr flow writes one).

If you have literally nothing (no branch name, no commit, no description),
**stop and ask the user** what the work is — a search seeded with empty
terms finds nothing and proves nothing.

## Step 1 — Resolve the repository

```bash
git remote get-url origin   # parse owner/repo from SSH or HTTPS form
gh auth status              # gh must be authenticated to search issues
```

- No origin remote, or the remote isn't a `github.com` URL → **stop** and
  ask. Don't search the wrong repo.
- `gh` not authenticated → **stop** and tell the user to authenticate.
  An unauthenticated search returning no rows is indistinguishable from
  "no related issue", so it must not be treated as a negative.

## Step 2 — Extract search keywords

Turn the inputs into query terms: lowercase, keep tokens of **4+ chars**,
drop English stopwords and repo-noise (e.g. `servicenow`, `sdk`, `pr`,
`issue`, `add`, `support`, `error` — words that appear in most titles and
add no discriminative signal). Keep order for longest-first queries.

```bash
# quick shell tokenizer
printf '%s' "$WORK_DESC" | tr '[:upper:]' '[:lower:]' | tr -cs 'a-z0-9' '\n' \
  | awk 'length($0) >= 4'
```

If the branch/context encodes an issue number (`123-` or `#123`), record
it — that is a *known* reference, which you should validate in step 3
rather than treat as an open question.

## Step 3 — Search open issues

GitHub search **ANDs every term**, so a long query returns zero rows even
when related issues exist. Walk a prefix ladder — all keywords first, then
progressively shorter prefixes (8, 4, then 3 of the highest-signal terms) —
and accumulate hits across every rung:

```bash
gh search issues --repo <owner>/<repo> --state open "<terms>" --limit 20
# or, for more control:
gh api search/issues -f q="repo:<owner>/<repo> is:issue is:open <terms>" \
  -f per_page=15 --jq '.items[] | [.number, .title, .state, .created_at] | @tsv'
```

- Exclude the obvious negatives: closed issues (work already done), PRs
  (`is:issue` filters those out), and any issue that is literally the one
  you already know about from a provided reference.
- If every rung returns nothing, that is a real negative — record
  `NO_RELATED_ISSUE` and stop.

## Step 4 — Score and rank candidates

Mirror the repo's own `issue-similarity-check` heuristic so this skill
agrees with how the maintainers judge "related":

- Lowercase each candidate's title and body.
- Score: **3 points per shared title keyword**, **1 per shared body
  keyword**, word-boundary matches only (so `test` cannot match
  `latest`).
- A candidate **qualifies** only if it shares **≥1 title keyword** and
  scores **≥ 5** — title overlap is what separates genuinely related
  issues from coincidental body overlap.

Then classify:

- **Strong match** — shares **≥ 2 title keywords** or scores **≥ 8**:
  `RELATED_ISSUE=<n>`.
- **Qualified but uncertain** — meets the qualify bar but not the strong
  bar, or there are several qualified candidates that could point at
  different work: `CANDIDATES=<n1,n2,...>`.
- **Nothing qualifies** — `NO_RELATED_ISSUE`.

Read the top candidates' actual titles/bodies before finalizing — a
handful of human reads beats scoring alone, and catches things like an
issue that is about a *different* verb of the same noun (a "docs" issue
named the same as a "feat" request). Don't let one fuzzy keyword carry a
link.

## Step 5 — Check whether the work is already in flight

For a strong (or user-selected) related issue, before declaring it safe
to link, check whether an **open PR already references it** — the work
may be in progress by someone else:

```bash
gh pr list --repo <owner>/<repo> --state open --search "<issue number or title>" --json number,title,headRefName
# or inspect linked PRs directly:
gh api repos/<owner>/<repo>/issues/<n>/timeline --jq '.[] | select(.event=="cross-referenced" and .source.issue.pull_request) | .source.issue.number'
```

If an open PR already closes the issue, include
`ALREADY_IN_PR=<m>` in the verdict. A fresh PR for the same issue is a
duplicate, not a fix — the caller should **stop** rather than ship.

## Step 6 — Declare the verdict

End the skill by stating exactly one of these, with the issue number(s)
and titles, so the caller can act without re-deriving anything:

- `RELATED_ISSUE=<n>` — strong match; caller links the PR (`Closes #<n>`,
  names the branch from it, and omits `no-issue-required`). Append
  `ALREADY_IN_PR=<m>` when an open PR already covers it.
- `CANDIDATES=<n1,n2,...>` — qualified but uncertain; caller asks the
  user which, if any, applies before proceeding.
- `NO_RELATED_ISSUE` — nothing open tracks the work; `no-issue-required`
  is an honest label.

## Notes on judgment calls

- **A false positive is worse than a miss.** Linking means the PR claims
  to resolve the issue. When in doubt, report candidates and let the user
  decide rather than declaring a strong link.
- **Same noun, different verb.** "Add license headers" and "Remove
  license headers" share keywords; only overlap in the *action* the PR
  performs counts as related.
- **`ALREADY_IN_PR` is a stop, not a prompt.** If the work is in flight,
  the useful outcome is joining that PR or closing the issue, not creating
  a parallel one.
- **Never fabricate a negative.** A failed `gh` call (rate limit,
  transient error) is not a "no related issue" result — surface the error
  and stop so the caller knows the check didn't run.