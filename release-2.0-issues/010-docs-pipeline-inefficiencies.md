# Docs pipeline: stale filter reference, redundant deploy rebuild

- **Priority:** P2
- **Raised by:** Senior DevOps
- **Area:** CI/CD (docs)

## Problem

In `.github/workflows/docs.yml`:

1. The `changes-docs` filter lists `.github/workflows/pages.yml` as a watched file, but
   the workflow file is `docs.yml` — `pages.yml` doesn't exist. Changes to `docs.yml`
   itself therefore don't mark docs as changed in the filter (the `on.push.paths`
   trigger catches it, but the job-level filter then reports no docs changes and jobs
   skip).
2. The `deploy` job downloads the built `site` artifact **and then runs
   `mkdocs gh-deploy`**, which rebuilds the site from scratch anyway — the artifact
   download and the earlier build/cache steps buy nothing for deploy. Either deploy the
   prebuilt artifact (e.g. `actions/deploy-pages` or `ghp-import site`) or drop the
   artifact plumbing.
3. The site cache keyed on `site-${{ github.sha }}` with `restore-keys: site-` can
   restore a *previous commit's* site and then skip the build entirely
   (`if: cache-hit != 'true'` only checks exact hit — partial restores do rebuild, so
   this one is mostly benign, but the cache adds little for an mkdocs build measured in
   seconds).

## Recommendation

Simplify: build once (`mkdocs build --strict`), upload artifact, deploy the artifact
with `actions/upload-pages-artifact` + `actions/deploy-pages` (the modern Pages flow —
the `id-token: write` permission is already declared but unused by `gh-deploy`). Fix or
remove the `pages.yml` filter entry and drop the cache.
