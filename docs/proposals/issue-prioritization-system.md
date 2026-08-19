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
- Not implementing anything against the GitHub Project board. The `project`
  scope was verified and the board's real fields/views were inspected while
  writing this draft (see "Current state" and "Mapping onto GitHub-native
  constructs" below), but no field/view/label was created or edited — this
  spec is still a design for the maintainer to apply by hand.

## Current state (as found 2026-07-25)

- **Labels already exist for priority**: `priority: low`, `priority: high`,
  `priority: urgent` (3 tiers). `priority: medium` does **not** exist yet —
  this spec adds it (see Design and Adoption plan below); it needs to be created
  before or during the retroactive scoring pass, not as a separate step.
- **`status:` labels** track triage/workflow status independently:
  `status: new`, `status: reviewed`, `status: blocked`, `status: duplicate`, `status: invalid`, `status: wontfix`.
  (Note: The original `state:` prefix was removed in Phase 1 of the project management redesign; `status:` is the canonical prefix.)
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
- **Project board state, verified 2026-07-25** (`gh project field-list 7
  --owner michaeldcanady`, `gh api graphql` for views; project #7, "ServiceNow
  SDK for Go," a private *user*-owned board, 41 items):
  - A **`Priority` single-select field already exists** with options `low` /
    `med` / `high` / `urgent` — i.e. the four label tiers this spec's rubric
    outputs are *already* mirrored as a board field, not something this spec
    needs to create. (This predates the spec being written; it just wasn't
    visible without the `project` scope.)
  - No `Priority score`-equivalent numeric field exists — the board has
    nothing today that expresses within-tier ordering.
  - No field named `Effort` exists, but a **`Size` single-select** field does
    (`XS` / `S` / `M` / `L` / `XL`), plus a separate plain **`Estimate`**
    field (untyped options, i.e. a number/text field, currently unused on any
    item in the sample pulled) and a **`Sprint`** iteration field. These are
    pre-existing sprint-planning scaffolding, not something this spec put
    there.
  - A **`Release`** single-select field already exists with options
    `No release` / `1.8.0` / `2.0.0`, separate from the native `Milestone`
    field this spec's Current-state section already described.
  - Existing views: `Backlog`, `Current sprint`, `My items`, `Needs Sprint`,
    `Needs Estimate`, `Sprints`, `Roadmap` — all sprint/estimation-flavored,
    none of them a triage or priority-sorted view.
  - `Status` (single-select) already includes a `Needs Triage` option,
    alongside workflow states (`Backlog - Ready`, `In progress`, `In review`,
    `Done`, `Backlog - Not Ready`, `Won't do`) and, oddly, two options that
    look copy-pasted from `Release` (`No release`, `Release 1.8.0`) — noted
    here as a pre-existing oddity on the board, not something this spec
    introduces or is scoped to clean up.

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
memory), the rubric's output stays a label. The original draft of this
section (written before the `project` scope was available) proposed adding
two brand-new Project v2 fields, `Priority score` and `Effort`, to carry
what the label alone can't. Now that the board itself has been inspected,
one of those two turns out to already partially exist under a different
name, and a third field (`Priority`) turns out to already duplicate the
label — see the reconciliation immediately below for what's actually built.

**Reconciled against the actual board (verified 2026-07-25, project #7 —
see "Current state"):** the plan above was written before the board could be
read; two of the three pieces already exist under names this spec didn't
anticipate, so the design changes as follows rather than adding net-new
fields on top of them:

- **`Priority` single-select (`low`/`med`/`high`/`urgent`) already exists on
  the board.** This is exactly the mirrored field the original draft argued
  *against* adding — it just predates this spec. Given it's already there
  (and, per `item-list`, already used on at least some items), reversing course
  and deleting it isn't worth the churn; instead, **the rubric's tier is the
  source of truth and this field is kept in sync with the `priority:` label**
  whenever the label is set. This does mean the "don't add a duplicate field"
  argument from the earlier draft is now moot — the duplication already
  existed independent of this spec — but the sync obligation is worth calling
  out explicitly so label and field don't drift.
- **`Priority score` (Number field) still needs to be created** — nothing on
  the board today expresses the raw 2–6 sum, so this part of the original
  design is unchanged.
- **`Effort` does not need to be created as a new field — reuse the existing
  `Size` single-select instead.** The board already has `Size`
  (`XS`/`S`/`M`/`L`/`XL`), which is finer-grained than the rubric's 1–3
  Effort scale but already exists and is presumably referenced by the
  pre-existing sprint views (`Needs Estimate`, `Sprints`). Adding a second,
  narrower `Effort` field alongside it would just be two fields answering
  "how big is this," one of which the rubric would ignore — the same
  field-proliferation problem the original draft was trying to avoid by
  *not* duplicating `Priority`. Map the rubric's Effort scale onto `Size` as
  1→`S`, 2→`M`, 3→`L`, and leave `XS`/`XL` available for cases finer than the
  rubric bothers to distinguish (a genuinely trivial fix vs. a genuinely
  sprawling one) rather than trying to force every issue into exactly three
  buckets. The separate `Estimate` and `Sprint` fields are sprint-planning
  scaffolding this repo doesn't currently use (no sprints per this spec's
  Non-goals) and are left untouched.

**Views to add** (none of the three below overlap with the board's existing
views — `Backlog`, `Current sprint`, `My items`, `Needs Sprint`,
`Needs Estimate`, `Sprints`, `Roadmap` are all sprint/estimation-flavored,
not triage- or priority-flavored, so they're additive, not a replacement):

- **"Triage" view** — filtered to `state: new`, grouped by nothing, sorted
  by `Priority score` descending. This is the "what needs a look" queue.
  (Note: `Status` already has a `Needs Triage` option that reads as a
  candidate alternative filter — using `state: new` instead keeps this
  spec's view aligned with the label taxonomy it's already built around,
  rather than introducing a second axis. Reconciling `Status`'s
  `Needs Triage` against the `state:` labels is a pre-existing overlap this
  spec doesn't resolve; flagged as a follow-up, not solved here.)
- **"Priority board" view** — grouped by the `Priority` field (the
  already-existing single-select, kept in sync per above), sorted by
  `Priority score` within each group. This is the steady-state working view.
- **"v2.0.0 board" view** — filtered to the `v2.0.0` **milestone** (the
  native GitHub field), not the board's `Release` single-select. `Release`
  already exists with a `2.0.0` option, but nothing in the sample pulled
  shows it populated, and this spec's rubric logic already keys off milestone
  membership (see "Applying it to this repo's context") — introducing
  `Release` as a second, currently-unpopulated proxy for the same thing would
  be one more field to keep in sync for no added benefit. Sorted by
  `Priority score`.

All three views and the `Priority score` field are still to be created by
hand — this spec only verified they don't already exist and don't collide
with what's there; it doesn't create them.

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

1. **`Status`'s `Needs Triage` option vs. the `state:` label taxonomy** — the
   board's `Status` field already has a `Needs Triage` option that overlaps in
   meaning with `state: new`. This spec's "Triage" view filters on the label,
   not `Status`, to stay consistent with the label-first design, but that
   means two fields can independently claim an issue is or isn't in triage.
   Whether to eventually retire one in favor of the other (or explicitly
   define which is authoritative) isn't resolved here.

Resolved since the previous draft (see commit history for this file):
whether to add a `priority: medium` tier (yes — see Design), how
`v2.0.0` milestone membership interacts with score (floors at `high`, does
not force `urgent` — see "Applying it to this repo's context," flagged
there as a judgment call the maintainer can override), what triggers
re-scoring (opportunistic, on touch — see "Applying it to this repo's
context"), and the Project board's actual fields/views now that the
`project` scope was granted (verified 2026-07-25 — the board already has a
`Priority` field to keep in sync, no existing `Effort`/`Priority score`
field, and `Size` should be reused instead of adding a new `Effort` field —
see "Current state" and "Mapping onto GitHub-native constructs").
