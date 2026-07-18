---
title: Why it's built this way
description: >-
  The SDK's three foundational design decisions, told in plain language —
  with the full ADRs one click deeper for the technical whys and hows.
---

# Why it's built this way

Every codebase has load-bearing decisions that look arbitrary until someone
explains them. This section is those explanations — short, plain-language
accounts of the three decisions the whole SDK rests on. Each page states the
decision, the reasons, and what it means for the code you write; each links
to its **Architecture Decision Record (ADR)** in `docs/adr/`, which is the
primary source when you want the deeper technical why and how, the
alternatives considered, and the trade-offs accepted.

Read these when a convention feels like friction — the answer to "why won't
review let me just…" is almost always one of these three.

## The decisions

| The question you were about to ask | The short answer | Deep dive |
| --- | --- | --- |
| *"Why is everything hand-written when Kiota has a generator?"* | ServiceNow publishes no reliable OpenAPI specs, and its response shapes need human judgment — so the repo hand-writes the surface on Kiota's runtime, and enforces generator-grade uniformity by convention instead. | [Why hand-written on Kiota?](design-hand-written-kiota.md) · [ADR 003](https://github.com/michaeldcanady/servicenow-sdk-go/blob/main/docs/adr/003-hand-written-on-kiota.md) |
| *"Why aren't models just structs?"* | A struct field can't tell "the instance sent an empty value" from "the instance didn't send this field" — a backing store can, and it dirty-tracks so `Patch` bodies only carry what you actually set. | [Why aren't models plain structs?](design-backed-models.md) · [ADR 002](https://github.com/michaeldcanady/servicenow-sdk-go/blob/main/docs/adr/002-backing-store-models.md) |
| *"Why can't I just `errors.New` here?"* | v1 had hundreds of similar-but-different error strings, which broke `errors.Is` for every consumer. Shared sentinels and one status-code mapping fixed that — a fresh `errors.New` reintroduces the bug. | [Why sentinel errors everywhere?](design-error-handling.mdx) · [ADR 001](https://github.com/michaeldcanady/servicenow-sdk-go/blob/main/docs/adr/001-error-standardization.md) |

## Making a decision of your own

Significant design changes follow the same lightweight norm that produced
these pages:

1. **Write the ADR** under `docs/adr/` in the repository — context, decision,
   consequences, alternatives considered. The ADR is the primary source and
   the artifact reviewers evaluate.
2. **Add (or update) the readable summary here** — a short page in the shape
   of the three above: the decision, the why, and "what to do in new code",
   linking back to the ADR.

If you're not sure a change rises to ADR level, open an issue and ask —
that conversation is cheaper than the rework either way.
