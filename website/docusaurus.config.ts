import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

// CI contract (see .github/workflows/docs-build.yml):
//   * DOCS_BASE_URL — path prefix the site is served under; defaults to the
//     production path. PR previews pass /servicenow-sdk-go/pr-preview/pr-<N>/.
//   * DOCS_ANALYTICS_DOMAIN — set ONLY for production builds. Absent/empty
//     means no Plausible script: local dev and PR previews stay out of stats.
const baseUrl = process.env.DOCS_BASE_URL || '/servicenow-sdk-go/';
const analyticsDomain = process.env.DOCS_ANALYTICS_DOMAIN;

// Release lines with an archive on GitHub Pages, newest first, such as
// "2.1,2.0". The stable deploy reads them from the X.Y/ folders on gh-pages;
// local builds and PR previews leave this unset and show no archive links.
// Each archive is a frozen build of its release tag, served under
// /servicenow-sdk-go/<line>/. Links use the production URL.
const archivedVersions = (process.env.DOCS_ARCHIVED_VERSIONS ?? '')
  .split(',')
  .map((line) => line.trim())
  .filter((line) => /^\d+\.\d+$/.test(line));
const siteUrl = 'https://michaeldcanady.github.io';
const archiveUrl = (line: string) => `${siteUrl}/servicenow-sdk-go/${line}/`;
const latestRelease = archivedVersions[0];

const config: Config = {
  title: 'ServiceNow SDK for Go',
  tagline: 'A fluent, type-safe Go client for ServiceNow REST APIs',
  favicon: 'img/favicon.png',

  future: {
    v4: true,
  },

  url: siteUrl,
  baseUrl,

  organizationName: 'michaeldcanady',
  projectName: 'servicenow-sdk-go',
  trailingSlash: false,

  onBrokenLinks: 'throw',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  markdown: {
    format: 'detect',
    mermaid: true,
    hooks: {
      onBrokenMarkdownLinks: 'throw',
    },
  },

  themes: ['@docusaurus/theme-mermaid'],

  clientModules: ['./src/clientModules/trackCodeCopy.ts'],

  scripts:
    process.env.NODE_ENV === 'production' && analyticsDomain
      ? [
          // Plausible analytics. Only loaded on production deploys: local
          // dev runs with NODE_ENV=development, and PR previews never set
          // DOCS_ANALYTICS_DOMAIN, so neither pollutes the stats.
          {
            src: 'https://plausible.io/js/script.js',
            defer: true,
            'data-domain': analyticsDomain,
          },
        ]
      : [],

  plugins: [
    // Allow importing Go snippet files as raw source (single source of truth
    // for all code examples; see src/components/GoSnippet and GoExample).
    function goRawSourcePlugin() {
      return {
        name: 'go-raw-source',
        configureWebpack() {
          return {
            module: {
              rules: [{test: /\.go$/, type: 'asset/source'}],
            },
          };
        },
      };
    },
    [
      require.resolve('@easyops-cn/docusaurus-search-local'),
      {
        hashed: true,
        indexBlog: false,
        docsRouteBasePath: '/',
        highlightSearchTermsOnTargetPage: true,
      },
    ],
  ],

  presets: [
    [
      'classic',
      {
        docs: {
          routeBasePath: '/',
          sidebarPath: './sidebars.ts',
          // Keep numeric prefixes (e.g. adrs/001-error-standardization) in doc
          // ids and URLs — ADR numbering is referenced as "ADR 003" everywhere.
          numberPrefixParser: false,
          // Released lines are archived builds of their release tags, not
          // copies in this tree (see contributing/release-branches), so every
          // page here is the live source.
          editUrl: 'https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/website/docs/',
          // "/" serves the docs from main. The version dropdown links to the
          // release archives in DOCS_ARCHIVED_VERSIONS.
          lastVersion: 'current',
          versions: {
            current: {label: 'main', banner: 'none'},
          },
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // main can document changes that haven't shipped yet, so point readers at
    // the newest release.
    announcementBar: latestRelease
      ? {
          id: `main-docs-${latestRelease}`,
          content: `These docs follow the main branch and can describe unreleased changes. For the latest release, see the <a href="${archiveUrl(latestRelease)}">v${latestRelease} docs</a>.`,
        }
      : undefined,
    navbar: {
      title: 'ServiceNow SDK for Go',
      logo: {
        alt: 'ServiceNow SDK for Go logo',
        src: 'img/logo.png',
      },
      items: [
        {type: 'docSidebar', sidebarId: 'userGuide', position: 'left', label: 'User Guide'},
        {type: 'docSidebar', sidebarId: 'apiReference', position: 'left', label: 'API Reference'},
        {type: 'docSidebar', sidebarId: 'contributing', position: 'left', label: 'Contributor Guide'},
        {
          type: 'docsVersionDropdown',
          position: 'right',
          dropdownItemsAfter: archivedVersions.map((line) => ({
            href: archiveUrl(line),
            label: `v${line}`,
          })),
        },
        {
          href: 'https://pkg.go.dev/github.com/michaeldcanady/servicenow-sdk-go',
          label: 'GoDoc',
          position: 'right',
        },
        {
          href: 'https://github.com/michaeldcanady/servicenow-sdk-go',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {label: 'Getting Started', to: '/getting-started'},
            {label: 'User Guide', to: '/user-guide'},
            {label: 'API Reference', to: '/apis'},
          ],
        },
        {
          title: 'Project',
          items: [
            {label: 'Releases', href: 'https://github.com/michaeldcanady/servicenow-sdk-go/releases'},
            {label: 'Roadmap', href: 'https://github.com/users/michaeldcanady/projects/7/views/9'},
            {label: 'Issues', href: 'https://github.com/michaeldcanady/servicenow-sdk-go/issues'},
          ],
        },
      ],
      copyright: `Community-driven project — not an official ServiceNow product.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ['go', 'bash'],
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
