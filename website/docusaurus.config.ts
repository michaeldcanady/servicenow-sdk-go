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
          editUrl:
            'https://github.com/michaeldcanady/servicenow-sdk-go/tree/main/website/',
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
