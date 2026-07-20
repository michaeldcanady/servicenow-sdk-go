# Support policy

This page states what Go versions and SDK versions are supported, so you can
plan upgrades with confidence.

## Go versions

The SDK supports [Go's official support window](https://go.dev/doc/devel/release#policy):
the latest two Go minor releases. Bumping the `go` directive in `go.mod` to track that window is
treated as a **minor** release, not a major one — it doesn't require a new
major version even though it can drop support for older Go toolchains.

## SDK versions

Only the latest major version is actively developed — new features and
non-critical fixes land there only.

Critical and security fixes are backported to the previous major version
(v1) for **approximately 6 months after v2.0.0 ships**. After that window,
v1 receives no further updates.

:::note
v1's sunset date will be published here once the v2.0.0 release date is
finalized.
:::

## What this means for you

- **On the latest major version:** you get new features, bug fixes, and
  security patches.
- **On the previous major version:** you get critical and security fixes
  only, for a limited time — plan your upgrade to the latest major version
  within that window.
- **Older than the previous major version:** no further updates; upgrade to
  a supported version.

## Next steps

- **[Migrating from v1 to v2](migrate-v1-to-v2.mdx):** Map v1 constructs to
  their v2 replacements.
- **[Getting Started](../getting-started.mdx):** Install and configure the
  latest version.
