# Issue prioritization system

## Problem

Issue triage in this repo (`michaeldcanady/servicenow-sdk-go`) is
inconsistent. There's already a `priority: low` / `priority: high` /
`priority: urgent` label set, but nothing defines *how* an issue earns one of
those labels — it's whatever the maintainer feels in the moment. That was
tolerable when the backlog was a handful of items, but it just grew sharply:
the `release-2.0-issues/` local markdown tracker was migrated into real
GitHub issues on 2026-07-25 (#555–#568), on top of whatever was already open,
specifically so everything lives in GitHub instead of a doc only the
maintainer reads. A pile of freshly-filed issues with no consistent priority
signal is exactly the situation a rubric needs to solve, and it needs to
solve it while the `v2.0.0` milestone (#13) is still in flight — the
in-flight release is itself a prioritization input (a v2.0-blocking issue and
a nice-to-have refactor aren't the same kind of urgent), not something to
design around.

This is a process/product change, not a code change — no request-builder,
error-handling, or model-layer convention is affected, so no ADR is drafted
alongside this spec (see CLAUDE.md's ADR-trigger list; issue triage isn't on
it).

## Goals

- A repeatable scoring rubric a single small-team maintainer can apply to an
  issue in under a minute, without needing a second reviewer to agree on the
  number.
- The rubric's *output* is the priority tiers that already exist (or will
  exist, once `priority: medium` is added — see Design) as labels
  (`priority: low` / `medium` / `high` / `urgent`) plus data the Project
  board can group and sort by — not a parallel taxonomy that competes with
  the labels.
- A plan for retroactively scoring the current open backlog (including
  #555–#568) once, rather than leaving old issues permanently un-scored.

## Non-goals

- Not designing a full Jira-style estimation process (story points, sprints,
  velocity tracking). This repo doesn't run sprints.
- Not committing to any automation (bot-scored issues, Actions-driven
  relabeling) in this round — automation is discussed as a future option,
  not built.
- Not changing the existing `type:`, `module:`, or `state:` label taxonomies.
  They're orthogonal to priority and already work.
- Not implementing anything against the GitHub Project board. This repo's
  `gh` token is currently missing the `project` scope (`gh project list`
  fails with "your authentication token is missing required scopes
  [read:project]"), so this spec's board-wiring section is a design for the
  maintainer to apply by hand (or after running
  `gh auth refresh -s project`) — not something this session could verify or
  execute even if it were in scope.

## Current state (as found 2026-07-25)

- **Labels already exist for priority**: `priority: low`, `priority: high`,
  `priority: urgent` (3 tiers). `priority: medium` does **not** exist yet —
  this spec adds it (see Design and Adoption plan below); it needs to be created
  before or during the retroactive scoring pass, not as a separate step.
- **`state:` labels** track triage/workflow status independently:
  `state: new`, `state: reviewed`, `state: in progress`, `state: blocked`.
- **`type:` labels** classify the kind of work: `type: bug`,
  `type: feature`, `type: refactor`, `type: documentation`, `type: devops`,
  `type: epic`, `type: test`.
- **`module:` labels** scope an issue to an API package (`module: table-api`,
  `module: core`, `module: cdm`, ~15 more) — useful for grouping but not a
  priority signal by itself.
- **One open milestone**: `v2.0.0` (#13, 3 open / 7 closed issues, ~14 issues
  total once #555–#568 are accounted for), described as "everything that must
  be decided or done before tagging v2.0.0." Milestone membership is itself a
  strong, already-modeled priority signal this rubric should lean on rather
  than duplicate — see "Applying it to this repo's context" for how it
  interacts with the score.
- **No custom issue fields** configured at the repo or org level
  (`list_issue_fields` returned empty) and **no GitHub Issue Types**
  configured (the issue-types endpoint 404s) — priority-relevant data today
  lives entirely in labels, not in typed fields.
- **Project board state is unknown** — the `gh` token lacks the `project`
  scope needed to even read board fields/views, let alone write them. This
  is a hard blocker for the board-wiring half of this spec until the
  maintainer runs `gh auth refresh -s project`.

## Design

### Rubric: Three dimensions, deliberately small

A 4th or 5th dimension (for example, "confidence," "reach") adds rigor real
product-triage processes need when many people are scoring many issues and
need to agree. That's not this repo's situation — one maintainer, occasional
outside contributors, no scoring disagreements to arbitrate. The rubric is
sized for "fast enough that it actually gets used every time," not "precise
enough to defend in a review meeting." Three dimensions, each 1–3:

1. **Impact** — how many consumers/how badly are they affected if this stays
   unaddressed?
   - 1 — cosmetic, docs-only, or affects an edge case / rarely-used module.
   - 2 — affects a commonly-used module (`tableapi`, `core`, `credentials`)
     or degrades DX without breaking correctness.
   - 3 — silent-failure/incorrect-behavior risk, a breaking change forced
     onto consumers, or blocks the `v2.0.0` milestone.
2. **Effort** — rough sizing, same granularity the maintainer already uses
   informally.
   - 1 — small, contained to one file/package, no design questions.
   - 2 — spans a few packages or needs a design decision but not an ADR.
   - 3 — cross-cutting, ADR-shaped, or touches the request-builder/model
     conventions CLAUDE.md tracks.
3. **Risk-of-delay** — what gets worse the longer this sits?
   - 1 — nothing; can sit indefinitely with no compounding cost.
   - 2 — accumulates tech debt or blocks a handful of other open issues.
   - 3 — the v2.0 release window makes this free *now* and expensive
     *forever after* (breaking changes, org/naming decisions — see ADR-010
     for exactly this kind of "free now, permanent later" reasoning) or it's
     already added to the `v2.0.0` milestone.

**Score = Impact + Risk-of-delay, with Effort as a tie-breaker, not an input
to the sum.** Effort deliberately isn't multiplied or added into the score:
folding effort into the same number as impact/risk lets a low-value
"quick win" outscore a real problem just because it's cheap, which inverts
the point of prioritizing. Instead:

- **Score 6** (Impact 3 + Risk 3 — both dimensions maxed) →
  `priority: urgent`
- **Score 5** (one dimension maxed, the other at 2) → `priority: high`
- **Score 4** (both dimensions at 2, or a 3+1 split) → `priority: medium`
- **Score 2–3** (both dimensions minimal, or only one mildly elevated) →
  `priority: low`
- Effort is used only to order issues *within* the same tier (cheaper first)
  and to flag when a high-impact/high-risk issue is also high-effort — that
  combination is worth calling out explicitly during triage (for example,
  via a triage comment) rather than silently encoding it in the label.

**Four tiers, derived by re-splitting the existing 2–6 range, not by
widening a dimension's scale.** The rubric's dimensions stay 1–3 exactly as
before — only the label mapping changes. Widening Impact or Risk-of-delay to
a 1–4 scale was considered and rejected: it would let a single dimension
independently exceed the other by more, which starts to matter for a
2-dimension additive sum in ways that would need the scoring guidance
itself to be rewritten (what does "Impact = 4" mean that "Impact = 3"
didn't?). Re-splitting the existing sum avoids that: the old 3-tier mapping
was `{5,6}→urgent`, `{3,4}→high`, `{2}→low`. The new 4-tier mapping peels the
top value off the old `urgent` band (6 stays `urgent`, 5 becomes its own
`high` tier) and relabels the old `high` band (3–4) as the new `medium`,
folding score 3 into `low` instead — because a score of 3 means at most one
dimension is even at "2," which reads closer to "mostly unremarkable" than
"medium." This keeps the bands roughly balanced against how many
Impact/Risk combinations produce each score (score 4 is reachable by three
different combinations, 3 and 5 by two each, 2 and 6 by exactly one), so no
single tier silently absorbs most of the backlog the way the old 2-value
`high` band risked doing.

This does mean the `priority: medium` label needs to be **created** in the
repo — it doesn't exist today (see "Current state"). Creating it folds into
the retroactive scoring pass in the Adoption plan below rather than being a
separate prerequisite step.

### Applying it to this repo's context

- **Milestone membership floors the tier at `high`; it doesn't force
  `urgent`.** Membership in `v2.0.0` still sets Risk-of-delay = 3 by
  definition (per the milestone's own description: "breaking changes are
  free here and expensive forever after"). Left purely to the score, an
  Impact-1 + Risk-3 milestone issue would land at score 4 (`medium`) under
  the new mapping — so this spec adds an explicit floor: **any issue in the
  `v2.0.0` milestone is set no lower than `priority: high`, regardless of
  what the raw score computes to**, while Impact still lets a milestone
  issue score `urgent` (Impact 3 + Risk 3 = 6) or stay at `high` (Impact 1
  or 2 + Risk 3 = 4 or 5, floored up to `high` if the raw score would say
  `medium`). Forcing every milestone issue straight to `urgent` was
  considered and rejected: with ~14 issues in `v2.0.0`, that would make
  every one of them indistinguishable at the top tier, which defeats the
  actual purpose of having a priority system at all — a board where triage
  can tell what to work on *next*, not just what's "in the release."
  Flooring at `high` instead still lets Impact differentiate a
  release-blocking correctness bug from a milestone-scoped documentation
  cleanup, while guaranteeing milestone work is never buried under
  non-milestone `low`/`medium` noise. **This is a judgment call, not a
  derived fact — flag it to the maintainer as something they can override**
  if in practice the milestone turns out to need its own `urgent`-only
  subset later.
- **Effort = 3 correlates with, but isn't identical to, "needs an ADR."**
  An Effort-3 issue that also touches a cross-cutting convention (see
  CLAUDE.md's ADR-trigger list) should get routed to a spec/ADR before
  scoring locks in a priority tier — scoring a design-shaped issue too early
  risks anchoring on an estimate made before the actual shape of the work is
  known.
- **A single maintainer means no reviewer-disagreement step.** Rubrics built
  for teams usually include a reconciliation step when two people's scores
  differ; that's cut entirely here, which is exactly why this rubric can
  stay to three dimensions instead of the five-plus a multi-scorer process
  usually needs.
- **Re-scoring is opportunistic, triggered by an issue being touched, not
  scheduled.** No new automated or calendar-driven re-triage process is
  introduced (per the decision to start fully manual — see Automation).
  Instead, a score/tier gets revisited whenever an issue is *already* being
  interacted with, so re-scoring rides along with work the maintainer is
  doing anyway rather than adding a new recurring chore:
  - the issue receives a new comment, is re-labeled, or is referenced from a
    PR — any of these is a natural moment to glance at whether the existing
    score/tier still holds;
  - the Effort estimate turns out to be off by 2+ tiers once real work
    starts (for example, scored Effort = 1 but it's turned into a
    multi-package change) — this is the one case worth explicitly flagging
    rather than leaving to "notice it eventually," since a wrong Effort
    estimate is exactly the kind of thing that should also prompt a second
    look at whether Impact/Risk were assessed correctly too;
  - an issue is reopened after being closed as won't-fix or duplicate —
    always re-score on reopen rather than restoring the stale label, since
    the circumstances that led to reopening usually mean *something* about
    the original assessment no longer holds.

### Mapping onto GitHub-native constructs

**Keep priority as a label, don't move it to a Project-only field.** Labels
are visible directly on the issue (in notifications, in `gh issue list`, in
search) without opening the board; a Project single-select field is
board-only. Since `priority: *` labels already exist (bar `medium`, added by
this spec) and are presumably referenced elsewhere (searches, muscle
memory), the rubric's output stays a label. **Do** add two Project v2 custom
fields that the label can't provide:

1. **`Priority score` (Number field)** — the raw 2–6 sum from the rubric.
   Lets the board sort *within* a tier instead of only grouping by it —
   three `priority: high` issues aren't equally urgent, and today there's no
   way to express that ordering anywhere.
2. **`Effort` (Single select: S / M / L, mirroring the 1/2/3 scale)** — kept
   separate from score per the "effort is a tie-breaker, not an input"
   design decision above. Having it as its own sortable field is what makes
   "cheap high-priority issues first" an actual board view instead of a
   mental note.

Do **not** add a Project-only `Priority` single-select duplicating the
label — GitHub Projects v2 can already group/filter by an issue's labels
directly, so a mirrored field would just be a second place the same value
has to be kept in sync, with no view it unlocks that grouping-by-label
doesn't already give.

**Views to add:**

- **"Triage" view** — filtered to `state: new`, grouped by nothing, sorted
  by `Priority score` descending. This is the "what needs a look" queue.
- **"Priority board" view** — grouped by the `priority:` label, sorted by
  `Priority score` within each group. This is the steady-state working view.
- **"v2.0.0 board" view** — filtered to the `v2.0.0` milestone, sorted by
  `Priority score`. Milestone-scoped work has its own working set distinct
  from the general backlog.

(All three are designs to apply once `project` scope is granted — see
Non-goals. Exact field/view creation commands aren't included here since
they're unverified against this repo's actual project number/layout.)

### Automation — flagged, not built

Two automation options exist for later, deliberately not pursued now:

1. **A GitHub Action that computes `Priority score` from label state** —
   feasible only if Impact/Risk inputs themselves become labels or issue-form
   fields (for example, an "Impact: 1/2/3" label set) the action can read. That's a
   second label taxonomy to maintain, which contradicts the "don't compete
   with existing labels" goal — not worth it until manual triage volume
   actually becomes a bottleneck.
2. **A required issue-template field** (`.github/ISSUE_TEMPLATE/*.yml`
   already has typed fields per CLAUDE.md's issue-writer skill) asking the
   *reporter* to self-score Impact — likely to be unreliable (reporters
   overstate impact) and shifts a maintainer judgment call onto whoever
   files the issue.

Recommendation: start fully manual (maintainer sets the label + two Project
fields during triage, using the rubric as a 30-second mental checklist, and
re-scoring opportunistically per the trigger list above). Revisit
automation only if triage volume grows enough that manual scoring becomes
the bottleneck — there's no evidence of that yet.

### Adoption plan

1. **Create the `priority: medium` label** as the first step of the
   retroactive pass below — it doesn't exist in the repo yet (see "Current
   state"), and there's no reason to make it a separate roll-out step ahead
   of the scoring pass that will immediately start applying it.
2. **Retroactive pass over #555–#568** (the just-migrated release-2.0
   issues) plus any other currently-open, unlabeled-for-priority issues —
   scored once, in one sitting, since they're already added to the
   milestone and context-fresh from today's migration. Apply the
   milestone-floor rule (see "Applying it to this repo's context") to the
   `v2.0.0` subset during this same pass.
3. **Prospective**: every newly filed issue gets a priority label + (once
   the Project fields exist) score/effort fields set during the same triage
   pass that currently applies `state: new` → `state: reviewed`. No new
   issue-template field is added (see Automation) — this is a triage-time
   action, not a filing-time requirement.
4. Document the rubric itself (dimensions, scale, bucket thresholds,
   milestone-floor rule, and re-scoring triggers) in `CONTRIBUTING.md` or a
   new `docs/TRIAGE.md` once this spec is approved — out of scope for this
   document, which is the design, not the contributor-facing writeup.

## Alternatives considered

- **5-dimension RICE-style rubric (Reach/Impact/Confidence/Effort/Score)** —
  rejected as over-built for a single-maintainer repo; the "Confidence"
  dimension in particular exists to hedge against multiple scorers'
  optimism bias, which doesn't apply here.
- **Numeric priority field only, drop the `priority:` labels** — rejected;
  breaks existing searches/filters/muscle-memory built around the labels,
  and a Project-only field is invisible off the board (see "Mapping onto
  GitHub-native constructs").
- **Impact × Effort × Risk (multiplicative) scoring** — rejected; a
  multiplicative score lets a 0-ish Effort term collapse the whole score
  regardless of Impact/Risk, which is the same "cheap wins outrank real
  problems" failure mode the additive-with-tiebreaker design avoids, just
  arrived at from the other direction.
- **Automate scoring from day one via issue-template fields** — rejected for
  now (see Automation section); adds a second taxonomy and shifts judgment
  onto reporters before there's evidence manual triage is a bottleneck.
- **Widening Impact and/or Risk-of-delay to a 1–4 scale to make room for a
  4th tier** — rejected in favor of re-splitting the existing 2–6 sum (see
  Design); widening a dimension changes what the rubric's own scoring
  guidance means per level and would need to be rewritten, whereas
  re-splitting only touches the label-mapping boundaries.
- **Forcing `priority: urgent` on every `v2.0.0`-milestone issue** —
  rejected (see "Applying it to this repo's context"); flattens ~14 issues
  into one indistinguishable tier, defeating the point of having tiers at
  all. Flooring at `high` instead was chosen so Impact can still
  differentiate within the milestone.
- **A scheduled/calendar-driven re-triage pass (for example, weekly)** — rejected in
  favor of an opportunistic, touch-triggered re-score; a single maintainer
  doesn't have the volume to justify a recurring ceremony, and tying
  re-scoring to actions already happening (comments, re-labels, PR
  references, reopens) means it costs nothing extra to remember.

## Open questions

1. **Project board number/layout** — this spec can't confirm the current
   project's fields or views exist because `gh`'s token lacks the `project`
   scope. Needs `gh auth refresh -s project` (or `read:project` for a
   read-only check first) before the board-wiring section can be executed
   or even verified against reality.

Resolved since the previous draft (see commit history for this file):
whether to add a `priority: medium` tier (yes — see Design), how
`v2.0.0` milestone membership interacts with score (floors at `high`, does
not force `urgent` — see "Applying it to this repo's context," flagged
there as a judgment call the maintainer can override), and what triggers
re-scoring (opportunistic, on touch — see "Applying it to this repo's
context").
