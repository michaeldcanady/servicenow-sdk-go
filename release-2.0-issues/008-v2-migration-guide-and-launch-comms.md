# v2 migration guide, deprecation policy, and launch communications

- **Priority:** P1 — product readiness
- **Raised by:** Senior Product Manager
- **Area:** Product / documentation

## Problem

v2 is a ground-up breaking rework (backing-store models, new fluent chaining, renamed
packages, new error contracts). There is currently no consumer-facing artifact that
answers: *"I'm on v1.12 — what do I do?"* Without it, existing users will pin v1
forever or churn, and the GitHub issue tracker becomes the migration guide.

There is also no stated support policy for the v1.x line, and the weekly preview
release channel's fate post-2.0 is undefined.

## Recommendation

1. **Migration guide** (`docs/migration-v2.md`, linked from Readme and the release
   notes): side-by-side v1→v2 examples for the top journeys — client construction,
   table CRUD, queries (`query.Field(...)`), attachments, pagination
   (`core.PageIterator`), error handling (`errors.Is` with `snerrors` sentinels),
   authentication via `credentials/`.
2. **v1 support statement**: e.g. "v1.x receives critical fixes only for N months; no
   new features." Put it in the Readme and the v2.0.0 release notes.
3. **Release notes**: hand-curate the v2.0.0 GitHub release body (release-please's
   generated changelog for a 90-commit squash will not tell the story).
4. **Preview channel**: decide whether the weekly preview continues (as `v2.x-preview`)
   or is retired at GA (see issue 011).
5. **Launch checklist**: verify pkg.go.dev renders the `/v2` docs, badges in Readme
   still resolve, and the docs site (mkdocs) reflects v2 before announcing.
