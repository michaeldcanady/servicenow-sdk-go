import React from 'react';
import CodeBlock from '@theme/CodeBlock';
import {extractRegion} from './region';
import {substitutePlaceholders, usePlaceholderValues} from './placeholders';

// Renders one named region of an imported Go snippet file as a code block.
//
//   import tablesGo from '@site/snippets/tables.go';
//   <GoSnippet src={tablesGo} region="table_get_fluent" />

export default function GoSnippet({
  src,
  region,
  language = 'go',
  title,
}: {
  src: string;
  region: string;
  language?: string;
  title?: string;
}): React.ReactElement {
  const values = usePlaceholderValues();
  return (
    <CodeBlock language={language} title={title}>
      {substitutePlaceholders(extractRegion(src, region), values)}
    </CodeBlock>
  );
}
