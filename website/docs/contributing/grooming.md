---
title: Issue grooming process
description: >-
  How triaged issues are promoted from Backlog to Ready — the checklist,
  workflow, and per-type guidance for making issues immediately actionable.
---

# Issue grooming process

This document defines how triaged issues are promoted from **Backlog** to **Ready** — meaning a contributor can pick them up and start coding immediately with no ambiguity.

Triage (covered in [Issue triage process](triage.md)) answers *"What is this and how important?"*
Grooming answers *"Is this ready to work on?"*

---

## When does grooming happen?

Grooming is not a scheduled ceremony. It's triggered by:

| Trigger                                | Action                                               |
| -------------------------------------- | ---------------------------------------------------- |
| Contributor asks "what can I work on?" | Groom the top-priority Backlog items first           |
| Weekly backlog scan                    | Review `Backlog` items, promote Ready where possible |
| Dependency resolves                    | Revisit issues that were blocked                     |
| New milestone created                  | Re-groom issues to assign scope                      |
| Issue receives new info                | Reporter responds to a question → re-evaluate        |

---

## The grooming checklist

Before an issue can leave `Backlog`, every box must be checked:

```
□ Acceptance criteria  — clear, testable "done" definition exists
□ Scope bounded        — not an epic in disguise; break into sub-issues if large
□ Dependencies clear   — no unresolved blockers
□ Effort estimated     — Size field set (S / M / L / XL)
□ Milestone assigned   — correct release milestone or "No release"
□ No open questions    — design questions resolved or explicitly deferred
```

Priority, type, and module labels are assumed set from triage. If missing, the issue belongs back in triage, not grooming.

---

## How to groom an issue

### Step 1: Check acceptance criteria

Can you write a test that proves this is done? If not, the criteria aren't clear enough.

- **For bugs:** reproduction steps must be verified, expected behavior documented
- **For features:** the "job to be done" from the template must translate to testable behavior
- **For refactors:** "no behavior change" must be verifiable (existing tests still pass)

If criteria are missing: comment on the issue requesting them. Keep in `Backlog`.

### Step 2: Check scope

Is this one piece of work, or several?

- If it touches more than ~3 files or spans multiple packages, consider whether it should be split into sub-issues linked under a parent epic.
- If it's tagged `type: epic`, it should already have sub-issues. If not, create them before grooming the parent.

### Step 3: Check for ADR-shaped work

Does this touch a cross-cutting convention (request-builder pattern, error handling, naming, pagination, nil-guards, backing-store models)?

- If **yes** and no ADR exists: draft the ADR first ([ADR collection](./adr/001-error-standardization.md)), keep in `Backlog` until accepted.
- If **yes** and an ADR already covers it: proceed.
- If **no**: proceed.

### Step 4: Set effort

Map the rubric's Effort score to the board's `Size` field:

| Effort | Size | Description                                            |
| ------ | ---- | ------------------------------------------------------ |
| 1      | S    | One file/package, no design decisions                  |
| 2      | M    | Few packages, minor design decision                    |
| 3      | L    | Cross-cutting or ADR-shaped (groom after ADR accepted) |
| —      | XS   | Trivial fix (typos, comment changes)                   |
| —      | XL   | Sprawling work that genuinely can't be split           |

When in doubt between two sizes, pick the smaller one. Over-scoping is worse than under-scoping.

### Step 5: Assign milestone

Every groomed issue needs a home:

- **Bug fix for current release** → current release milestone
- **New capability** → next minor milestone
- **Tech debt / DX improvement** → next minor or "No release"
- **No clear timeline** → "No release" (revisit quarterly)

If no forward-looking milestone exists yet, create one before grooming.

### Step 6: Verify no open questions

Scan the issue comments. Are there unanswered questions about approach, scope, or behavior?

- If **yes**: resolve them (comment with decision) or explicitly defer (comment "deferred to implementation — recording approach here: ..."). Don't leave ambiguity.
- If **no**: promote to Ready.

---

## Promotion

Once all checklist items are met:

1. **On the project board:** move Status from `Backlog` → the issue is now Ready
2. **Leave `status: reviewed` label** — this is the label equivalent of `Backlog`
3. **Optionally comment:** "Groomed and ready for pickup" with a summary of what "done" means

A contributor sees a Ready issue and knows: it has criteria, it's scoped, it's estimated, it has a milestone. They can start immediately.

---

## Grooming by issue type

### Bugs

Bugs need extra grooming attention:

- **Reproduction verified?** If the repro code in the template doesn't actually reproduce, comment and ask for clarification before grooming.
- **Root cause identified?** Not always possible, but if you know *why* it's broken, note it in a comment — saves the implementer significant time.
- **Regression or new?** If regression, link to the PR that introduced it.

### Features

Features need scope discipline:

- **"Job to be done" testable?** The template uses "When X, I want Y, so I can Z." Can you write a test for Y? If not, break it down further.
- **API shape proposed?** If the feature involves new request builders or models, sketch the intended Go API in a comment. This is not the implementation — it's the acceptance criteria.
- **New module?** If `module: new`, a full API surface discovery is needed before grooming. Use the [Add a new API module](add-api-module.md) playbook.

### Epics

Epics are grooming-resistant by nature. They need decomposition before they can be Ready:

- **Has sub-issues?** If not, create them. An epic without sub-issues is just a large issue pretending to be planned.
- **Sub-issues groomed individually?** Each sub-issue goes through this same checklist. The epic itself stays in `Backlog` until enough sub-issues are Ready to start.
- **No milestone on the epic itself.** Epics span multiple milestones — milestones belong on the sub-issues, not the epic.

### Decomposing stories into tasks

When a story (feature/bug) needs to be broken into tasks:

- **Use GitHub sub-issues.** Link tasks as sub-issues of the story via the GitHub UI or `gh issue edit <task> --add-parent <story>`.
- **Story stays open.** The story represents user value — it closes when the feature is complete, not when tasks are created.
- **Tasks are implementation details.** They close when the work is done. The story stays open until all tasks close.

### Work item formats

Different work item types use different formats. Don't mix them up:

**Stories** (user-facing value):
- "Job to be done" format: When [situation], I want [action], so I can [outcome]
- Acceptance criteria
- Describes value to the user

**Tasks** (technical work):
- Description: What needs to be done (concise, technical)
- Exit criteria: When is it done
- Story: Where it came from (parent issue link)

**Spikes** (research/investigation):
- Description: What uncertainty needs to be resolved
- Timebox: How long to spend (e.g., 1 day, 2 days)
- Background: Why this spike is needed
- Acceptance criteria: What the spike must produce (decision, recommendation, PoC)
- Output: ADR, documentation, or follow-up issues — not code

### Refactors / tech debt

- **"No behavior change" verifiable?** Link to existing tests that prove it. If none exist, add tests first (that's a separate issue).
- **Deprioritized vs. forgotten?** If it's `priority: low` with no milestone, it's effectively forgotten. Either assign a milestone or close it with a "won't fix" and a note that it can be revisited.

---

## Common mistakes

1. **Grooming without criteria.** "I know what this means" is not acceptance criteria. Write it down.
2. **Grooming epics as single issues.** If it's bigger than ~3 files, split it.
3. **Grooming Effort-3 issues before the ADR.** Design-shaped work needs a design document first.
4. **Assigning milestone without scope.** A milestone with 30 issues and no target date is a wish list, not a plan.
5. **Skipping Size because "it's obvious."** The Size field is for the *next* person, not you. They don't have your context.

---

## See also

- [Issue triage process](triage.md)
- [Conventions reference](conventions.md)
