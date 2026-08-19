# Project Management System Audit & Recommendation

## Executive Summary

The servicenow-sdk-go repository has **multiple overlapping project management systems** that create confusion and fragmentation. This audit identifies the gaps and proposes a unified system.

---

## Current State Audit

### What Exists Today

| System                   | Location                     | Status                               |
| ------------------------ | ---------------------------- | ------------------------------------ |
| **Specs**                | `specs/001-007/`             | 7 specs, varying completion          |
| **ADRs**                 | `docs/adr/001-010`           | 10 decisions documented              |
| **GitHub Issues**        | Issues #555-#568 + others    | Recently migrated from local tracker |
| **GitHub Project Board** | Project #7                   | 41 items, sprint-focused views       |
| **Priority Proposal**    | `docs/proposals/`            | Designed but not implemented         |
| **CLAUDE.md**            | Root                         | Detailed architecture context        |
| **Contributing Guide**   | `website/docs/contributing/` | Comprehensive                        |

### Critical Gaps Identified

#### 1. **No Work Item Linkage**
- Specs have `tasks.md` with checkboxes, but these aren't linked to GitHub issues
- No traceability from spec → issue → PR → release
- Example: `specs/002-v2-release-prep/tasks.md` has 25 tasks, none linked to issues

#### 2. **Dual Priority Systems**
- Labels: `priority: low/high/urgent` (missing `medium`)
- Project field: `Priority` (single-select, already exists)
- Proposal: Impact + Risk-of-delay rubric (not implemented)
- **Result**: Inconsistent prioritization, no scoring

#### 3. **Status Confusion**
- `state:` labels: `new`, `reviewed`, `in progress`, `blocked`
- `Status` field: `Needs Triage`, `Backlog - Ready`, `In progress`, `In review`, `Done`, `Backlog - Not Ready`, `Won't do`
- **Result**: Two overlapping taxonomies, unclear which is authoritative

#### 4. **No Sprint/Iteration Cadence**
- `Sprint` iteration field exists on project board
- `Size` field exists (XS/S/M/L/XL)
- No actual sprint planning or time-boxed iterations
- **Result**: Backlog-only, no delivery rhythm

#### 5. **Missing Post-v2 Roadmap**
- Only one milestone: `v2.0.0` (now released)
- No `v2.1.0`, `v3.0.0`, or feature milestones
- **Result**: No forward-looking planning visibility

#### 6. **Spec System Underutilized**
- Specs are comprehensive (user stories, acceptance criteria, tasks)
- But: No status tracking, no completion %, no linkage to implementation
- Example: `specs/001-prerelease-workflow` is "Draft" status but appears implemented

---

## Recommended New System

### Principles
1. **Single source of truth**: GitHub Issues as the work item hub
2. **Lightweight specs**: Only for cross-cutting/architectural work
3. **Clear ownership**: Every issue has an assignee
4. **Time-boxed delivery**: Milestones with target dates
5. **Automated status**: Reduce manual labeling burden

### Proposed Structure

#### 1. Issue Taxonomy (Simplified)

**Remove** `state:` labels entirely. Use GitHub Project `Status` field only:

| Status        | Meaning                                            |
| ------------- | -------------------------------------------------- |
| `Backlog`     | Triaged, not yet prioritized for work              |
| `Ready`       | Prioritized, dependencies resolved, ready to start |
| `In Progress` | Actively being worked                              |
| `In Review`   | PR open, awaiting review                           |
| `Done`        | Merged and verified                                |

**Keep** `type:` labels (bug, feature, refactor, documentation, devops, epic, test)

**Keep** `module:` labels (table-api, core, credentials, etc.)

**Implement** priority rubric from proposal:
- `priority: urgent` (Score 6)
- `priority: high` (Score 5)
- `priority: medium` (Score 4) - **create this**
- `priority: low` (Score 2-3)

#### 2. Milestone Structure

Replace single `v2.0.0` with quarterly milestones:

```
v2.1.0 (Q4 2026) - API surface expansion
v2.2.0 (Q1 2027) - Developer experience
v3.0.0 (Q3 2027) - Next major (if needed)
```

Each milestone has:
- Target date
- Description of scope
- Issues assigned via `milestone:` field

#### 3. Spec System (Reduced)

**Keep specs only for**:
- Cross-cutting architectural changes (error handling, model patterns)
- New API module blueprints
- Breaking change proposals

**Don't create specs for**:
- Bug fixes
- Single-file changes
- Documentation updates

**Spec lifecycle**:
```
Draft → Approved → In Progress → Complete
```

#### 4. Project Board Views

Replace current sprint-focused views:

| View               | Filter                 | Sort                 | Purpose               |
| ------------------ | ---------------------- | -------------------- | --------------------- |
| **Triage**         | `Status = Backlog`     | Priority score desc  | What needs a look     |
| **Ready to Start** | `Status = Ready`       | Priority score desc  | What to work on next  |
| **In Progress**    | `Status = In Progress` | Assignee             | Active work           |
| **In Review**      | `Status = In Review`   | Created asc          | PRs awaiting review   |
| **By Module**      | All                    | Group by `module:`   | Module health         |
| **By Priority**    | All                    | Group by `priority:` | Priority distribution |
| **Release Prep**   | Milestone = current    | Status               | Release readiness     |

#### 5. Automation Opportunities

| Automation                                   | Trigger       | Action               |
| -------------------------------------------- | ------------- | -------------------- |
| Auto-assign `state: new` → `Status: Backlog` | Issue created | Set initial status   |
| Auto-set `Status: In Review`                 | PR opened     | Move issue status    |
| Auto-set `Status: Done`                      | PR merged     | Close issue          |
| Priority score calculation                   | Label changed | Update project field |
| Milestone progress                           | Issue closed  | Update milestone %   |

#### 6. Documentation Updates

Update these files to reflect new system:

| File                         | Changes                                    |
| ---------------------------- | ------------------------------------------ |
| `CONTRIBUTING.md`            | Add priority rubric, issue lifecycle       |
| `CLAUDE.md`                  | Update subagent conventions for new labels |
| `website/docs/contributing/` | Add project management guide               |
| `docs/TRIAGE.md`             | **New** - triage process documentation     |

---

## Implementation Plan

### Phase 1: Cleanup (Week 1)
- [ ] Close/archive completed specs (001, 002, 003, 004, 005, 006)
- [ ] Create `priority: medium` label
- [ ] Retroactively score all open issues using rubric
- [ ] Remove duplicate `state:` labels from issues (keep `Status` field)

### Phase 2: Structure (Week 2)
- [ ] Create v2.1.0 milestone with target date
- [ ] Create new project board views (Triage, Ready, In Progress, In Review)
- [ ] Migrate v2.0.0 milestone issues to v2.1.0 or close
- [ ] Document triage process in `docs/TRIAGE.md`

### Phase 3: Automation (Week 3)
- [ ] GitHub Action: Auto-set status on PR open/merge
- [ ] GitHub Action: Sync priority label ↔ project field
- [ ] Update CLAUDE.md subagent conventions

### Phase 4: Adoption (Ongoing)
- [ ] Triage all new issues within 48 hours
- [ ] Weekly review of `Backlog` → `Ready` promotion
- [ ] Monthly milestone progress review

---

## Metrics for Success

| Metric                            | Current | Target             |
| --------------------------------- | ------- | ------------------ |
| Issues without priority label     | ~50%    | <10%               |
| Issues without assignee           | ~80%    | <20%               |
| Average time from `New` → `Ready` | Unknown | <7 days            |
| Milestone completion rate         | N/A     | >80% on-time       |
| Spec-to-issue linkage             | 0%      | 100% for new specs |

---

## Next Steps

1. **Review this audit** with maintainer
2. **Prioritize which phases** to implement first
3. **Create GitHub Issue** to track this project management improvement
4. **Assign owner** for implementation

---

## Open Questions

1. Do you want to implement this full system, or start with a subset?
2. Should we keep the spec system for future work, or retire it entirely?
3. What's your preferred sprint cadence (if any) - weekly, bi-weekly, monthly, or none?
