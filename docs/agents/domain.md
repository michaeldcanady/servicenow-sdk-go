# Domain Docs

How the engineering skills should consume this repo's domain documentation when exploring the codebase.

## Before exploring, read these

- **`GLOSSARY.md`** at the repo root.
- **`website/docs/contributing/adrs/`**: read the ADRs that touch the area you're about to work in. The catalog index is `index.md`; the plain-language overview of the foundational decisions is `../design-decisions.md`.
- **`docs/rules/`**: the agent-facing rule files distilled from the ADRs (`R-{AREA}-{NNN}.md`). Each carries RFC-2119 MUST/MUST NOT lines and a runnable `## Verify` command. Consult the rule files for the areas your change touches; run their `## Verify` commands before finalizing. A rule file and its ADR share the same number.

If any of these files don't exist, **proceed silently**. Don't flag their absence; don't suggest creating them upfront. The `/domain-modeling` skill (reached via `/grill-with-docs` and `/improve-codebase-architecture`) creates them lazily when terms or decisions actually get resolved.

## File structure

Single-context repo:

```
/
├── GLOSSARY.md
├── docs/rules/                 # agent rule files (RFC-2119 distillations of the ADRs)
│   ├── R-ERR-001.md
│   └── R-MODEL-002.md
└── website/docs/contributing/
    ├── design-decisions.md
    └── adrs/
        ├── index.md
        ├── 001-error-standardization.md
        └── 002-backing-store-models.md
```

Each `docs/rules/R-{AREA}-{NNN}.md` pairs with the ADR of the same number (`NNN`). The `NNN` is the shared key; `AREA` scopes the concern (ERR, MODEL, BUILD, PAGE, NAME, REL, AUTH, SURF, TBL, ATT). Rule files live outside `website/docs/` so Docusaurus does not publish them — they are for agents reading the filesystem, not the human docs site.

## Use the glossary's vocabulary

When your output names a domain concept (in an issue title, a refactor proposal, a hypothesis, a test name), use the term as defined in `GLOSSARY.md`. Don't drift to synonyms the glossary explicitly avoids.

If the concept you need isn't in the glossary yet, that's a signal: either you're inventing language the project doesn't use (reconsider) or there's a real gap (note it for `/domain-modeling`).

## Flag ADR conflicts

If your output contradicts an existing ADR, surface it explicitly rather than silently overriding:

> _Contradicts ADR-002 (backing-store models), but worth reopening because…_
