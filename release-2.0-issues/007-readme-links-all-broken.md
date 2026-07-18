# Readme API table links point to directories that no longer exist

- **Priority:** P1 — launch-facing documentation
- **Raised by:** Senior Product Manager
- **Area:** Documentation

## Problem

Every entry in the "Supported Service-Now APIs" table in `Readme.md` links to the old
hyphenated directory layout (`account-api`, `actsub-api`, `table-api`,
`documents-api`, ...). Those directories were renamed to the current unhyphenated form
(`accountapi`, `actsubapi`, `tableapi`, ...), so **all 15 links 404**. The table is
also missing modules added since (`statsapi`, and possibly others), and the links point
at `tree/main/...` even though 2.0 content is what will ship.

Verified:

```
$ ls account-api actsub-api table-api
ls: cannot access 'account-api': No such file or directory  (all three)
```

## Recommendation

1. Regenerate the table from the actual `*api` directories; link to each package's
   `Readme.md`.
2. Add the missing modules and remove any that no longer ship in v2.
3. Update the quick-start snippet to the v2 fluent API
   (`client.Now().Table("incident")...`) and — once issue 003 lands — the `/v2` import
   path.
4. Add a CI-friendly link check (e.g. `lychee` or `markdown-link-check` on `Readme.md`
   and `docs/`) so this class of rot is caught automatically.
