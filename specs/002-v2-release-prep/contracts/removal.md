# API Contract Changes: Removals

This document lists the major interface and function removals for v2.0.

## Authentication Providers
- `BasicAuthenticationProvider` is now the standard for username/password authentication.
- `ROPCCredential` is the standard for OAuth2 token-based authentication.

## Service Clients
- `NewServiceNowServiceClient` is the only entry point. `NewServiceNowClient` and other variants are removed.

## Request Builders
- All `*2` and `*3` versions of methods (e.g., `Get2`, `Post4`) will be renamed to their base names (e.g., `Get`, `Post`) OR their older counterparts will be deleted, and the new ones will become the primary API.
- For v2.0, we will prioritize naming consistency.

## Query Builders
- `query2` fluent API is the only supported query builder.
- `query` package is deleted.
