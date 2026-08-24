import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

const config: Config = {
  title: 'ServiceNow SDK for Go',
  tagline: 'A fluent, type-safe Go client for ServiceNow REST APIs',
  favicon: 'img/favicon.png',

  future: {
    v4: true,
  },

  url: 'https://michaeldcanady.github.io',
  // Overridden by CI for PR previews, which deploy under
  // /servicenow-sdk-go/pr-preview/pr-<N>/ (see .github/workflows/docs.yml).
  baseUrl: process.env.DOCS_BASE_URL ?? '/servicenow-sdk-go/',

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
    process.env.NODE_ENV === 'production' && !process.env.DOCS_BASE_URL
      ? [
          // Plausible analytics. Only loaded on production deploys: local
          // dev runs with NODE_ENV=development, and PR previews inject
          // DOCS_BASE_URL, so neither pollutes the stats.
          {
            src: 'https://plausible.io/js/script.js',
            defer: true,
            'data-domain': 'michaeldcanady.github.io',
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
          // Only "main" offers edit links. Frozen versioned copies are fixed
          // in docs/ first, then backported (see contributing/release-branches),
          // so pointing "Edit this page" at a snapshot would invite edits in
          // the wrong direction.
          editUrl: ({version, docPath}) =>
            version === 'current'
              ? `https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/website/docs/${docPath}`
              : undefined,
          // "/" serves the docs from main; released doc lines live under
          // their version prefix (e.g. /2.0/...) via the navbar dropdown.
          lastVersion: 'current',
          versions: {
            // The unreleased banner points readers at the newest released
            // line (/2.0/) instead of silently documenting unreleased code.
            current: {label: 'main', banner: 'unreleased'},
            '2.0': {label: 'v2.0', banner: 'none'},
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
