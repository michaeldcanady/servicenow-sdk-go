// Builds a frozen archive of the docs site from a release tag.
//
// The docs-archive workflow copies this file into the website/ directory of a
// tag checkout and runs `docusaurus build --config docusaurus.archive.config.ts`.
// It wraps the tag's own docusaurus.config.ts, so the archive renders with the
// components, sidebars, and Go snippets that shipped with that release. It
// changes only what makes the build an archive:
//
//   * Only the tag's docs/ directory is built, labeled with the release line.
//     Any versioned_docs/ copies at the tag are skipped.
//   * "Edit this page" links are removed, because an archive is never edited.
//   * The version dropdown is replaced with a link to the main docs site,
//     whose dropdown lists every release.
//   * An announcement bar says the reader is on archived docs. It doesn't
//     name a latest release, so it stays true after newer releases ship.
//
// Environment:
//   * DOCS_ARCHIVE_VERSION: the release line, such as 2.1. Required.
//   * DOCS_BASE_URL: the path prefix the archive is served under, such as
//     /servicenow-sdk-go/2.1/. The tag's config reads it.
//   * DOCS_ANALYTICS_DOMAIN: optional; the tag's config reads it.
import type {Config} from '@docusaurus/types';
import base from './docusaurus.config';

const version = process.env.DOCS_ARCHIVE_VERSION;
if (!version || !/^\d+\.\d+$/.test(version)) {
  throw new Error(`DOCS_ARCHIVE_VERSION must look like 2.1; got "${version ?? ''}".`);
}

const mainSiteUrl = `${base.url}/${base.projectName}/`;

type PresetEntry = [string, Record<string, any>];

const presets = (base.presets ?? []).map((preset) => {
  if (!Array.isArray(preset) || preset[0] !== 'classic') {
    return preset;
  }
  const [name, options] = preset as PresetEntry;
  return [
    name,
    {
      ...options,
      docs: {
        ...options.docs,
        editUrl: undefined,
        lastVersion: 'current',
        onlyIncludeVersions: ['current'],
        versions: {current: {label: `v${version}`, banner: 'none'}},
      },
    },
  ];
});

const themeConfig = (base.themeConfig ?? {}) as Record<string, any>;
const navbar = themeConfig.navbar ?? {};
const navbarItems = (navbar.items ?? []).map((item: Record<string, any>) =>
  item.type === 'docsVersionDropdown'
    ? {href: mainSiteUrl, label: `v${version} (archived)`, position: item.position ?? 'right'}
    : item,
);

const config: Config = {
  ...base,
  presets: presets as Config['presets'],
  themeConfig: {
    ...themeConfig,
    navbar: {...navbar, items: navbarItems},
    announcementBar: {
      id: `archived-${version}`,
      content: `You're reading the archived docs for v${version}. For other releases, or for the docs on the main branch, go to the <a href="${mainSiteUrl}">main docs site</a>.`,
      isCloseable: false,
    },
  },
};

export default config;
