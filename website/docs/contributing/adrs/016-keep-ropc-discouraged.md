---
title: 'ADR 016: Keep the ROPC grant, but steer callers away from it'
description: >-
  The resource owner password credentials grant stays available for legacy
  ServiceNow setups, with no removal plan and guidance against it in the godoc
  and the authentication guide.
---

# ADR 016: Keep the ROPC grant, but steer callers away from it

## Status

Accepted

## Context

OAuth 2.1 and RFC 9700 deprecate the resource owner password credentials
(ROPC) grant because it hands the SDK the user's password and defeats
multi-factor authentication and federation. ServiceNow still supports the
grant, and some legacy setups have no substitute: no interactive browser for
the authorization code flow, and no client secret that the client credentials
grant requires.

Alternatives considered:

1. **Remove the grant** — rejected: it breaks working legacy setups with no
   migration target.
2. **Mark it `// Deprecated:` and plan removal** — rejected: the marker reads
   as "this will break", which is not true; the SDK has no removal plan and
   should not imply one.
3. **Keep it without comment** — rejected: callers miss the security trade-off.
4. **Keep it and recommend against it in the docs** — chosen.

## Decision

`NewROPCProvider` and its credential stay. They carry no `// Deprecated:`
godoc marker and no removal plan. Two places recommend against the grant: the
godoc on `NewROPCProvider` and the authentication guide. Both point callers to
the client credentials grant for machine identity and the authorization code
flow for user identity.

## Consequences

- **Pros:** legacy setups keep working; new callers see the trade-off and a
  better grant before they choose; the SDK does not promise a removal it does
  not intend.
- **Cons:** the grant remains a supported path a caller can reach; the guidance
  lives in prose, so nothing stops a caller who ignores it.
