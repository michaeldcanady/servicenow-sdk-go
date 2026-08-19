# Documentation site

The public documentation site for the ServiceNow SDK for Go, built with
[Docusaurus](https://docusaurus.io/) and deployed to GitHub Pages by
`.github/workflows/docs.yml`.

## Layout

- `docs/` — the site's pages (Markdown/MDX). Sidebar structure is defined in
  `sidebars.ts`.
- `snippets/` — Go source files that are the single source of truth for every
  code example. Regions are marked with `// [START name]` / `// [END name]`
  comments and rendered through the `GoSnippet` (one region) and `GoExample`
  (assembled program from a template) components in `src/components/`.
- `doc-templates/` — authoring templates for new pages (not published).

## Commands

```bash
npm ci        # install (or: just setup-docs)
npm start     # live-reload dev server (or: just serve-docs)
npm run build # production build into build/ (or: just generate-docs)
```

The build fails on broken internal links (`onBrokenLinks: 'throw'`), so run
`npm run build` before pushing doc changes.
