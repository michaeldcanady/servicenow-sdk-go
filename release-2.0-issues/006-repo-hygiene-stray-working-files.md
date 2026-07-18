# Repo hygiene: development scratch files committed at the root

- **Priority:** P1 — first-impression / professionalism at launch
- **Raised by:** Senior Product Manager (with Principal Engineer)
- **Area:** Repository quality

## Problem

The repo root — the first thing every evaluating engineer sees on GitHub and the
directory listing pkg.go.dev links to — contains committed working artifacts:

| File | What it is |
| --- | --- |
| `coverage.html`, `coverage.out` | generated coverage output |
| `files_to_fix.txt` | a scratch worklist from a past refactor |
| `fix_error_mappings.py` | a one-off migration script (references `files_to_fix.txt`) |
| `GEMINI.md` | assistant context file, duplicative of CLAUDE.md |
| `placeholder-plugin.yaml` | placeholder |
| `specs/001…006` | internal planning docs for the v2 effort |
| `mocking.go` at root | test helper living in the public root package |

`spec/` (ServiceNow OpenAPI specs) is arguably legitimate reference material but should
be distinguished from `specs/` (internal planning) — the near-identical names invite
confusion.

## Recommendation

1. Delete `coverage.html`, `coverage.out`, `files_to_fix.txt`, `fix_error_mappings.py`,
   `placeholder-plugin.yaml`; add `coverage.*` to `.gitignore`.
2. Decide on `GEMINI.md`: delete or fold into CLAUDE.md.
3. Move `specs/` planning docs to `docs/design/` or delete if historical; consider
   renaming `spec/` → `api-specs/` for clarity.
4. Audit root-package exported surface: `mocking.go` in `package servicenowsdkgo`
   ships mock helpers to every consumer — move under `internal/mocking` or a clearly
   named `snmock` package if it's meant to be public.
