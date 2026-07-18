import type {SidebarsConfig} from '@docusaurus/plugin-content-docs';

// Mirrors the former mkdocs.yml nav, split into one sidebar per navbar tab.

const sidebars: SidebarsConfig = {
  userGuide: [
    'index',
    'getting-started',
    {
      type: 'category',
      label: 'Authentication',
      link: {type: 'doc', id: 'user-guide/authentication/index'},
      items: [
        'user-guide/authentication/ropc',
        'user-guide/authentication/auth-code-private',
        'user-guide/authentication/auth-code-public',
        'user-guide/authentication/basic',
        'user-guide/authentication/client-credentials',
        'user-guide/authentication/jwt-token',
      ],
    },
    'user-guide/tables',
    {
      type: 'category',
      label: 'Attachments',
      link: {type: 'doc', id: 'user-guide/attachments/index'},
      items: [
        'user-guide/attachments/upload-from-disk',
        'user-guide/attachments/download-record-attachments',
        'user-guide/attachments/most-recent-attachment',
        'user-guide/attachments/delete-aging-attachments',
      ],
    },
    'user-guide/batch',
    'user-guide/pagination',
    'user-guide/preview-features',
    'user-guide/query-builder',
  ],
  apiReference: [
    'apis/index',
    {
      type: 'category',
      label: 'Table API',
      link: {type: 'doc', id: 'apis/tables/index'},
      items: [
        'apis/tables/list',
        'apis/tables/create',
        'apis/tables/get',
        'apis/tables/update',
        'apis/tables/delete',
      ],
    },
    {
      type: 'category',
      label: 'Attachment API',
      link: {type: 'doc', id: 'apis/attachment/index'},
      items: [
        {
          type: 'category',
          label: 'File',
          link: {type: 'doc', id: 'apis/attachment/file/index'},
          items: ['apis/attachment/file/get', 'apis/attachment/file/create'],
        },
        {
          type: 'category',
          label: 'Upload',
          link: {type: 'doc', id: 'apis/attachment/upload/index'},
          items: ['apis/attachment/upload/create'],
        },
        'apis/attachment/list',
        'apis/attachment/get',
        'apis/attachment/delete',
      ],
    },
    {
      type: 'category',
      label: 'Batch API',
      link: {type: 'doc', id: 'apis/batch/index'},
      items: ['apis/batch/create'],
    },
  ],
  contributing: [
    'contributing/index',
    'contributing/setup',
    'contributing/architecture',
    'contributing/testing',
  ],
};

export default sidebars;
