import React from 'react';
import CodeBlock from '@theme/CodeBlock';
import {extractRegion} from '../GoSnippet/region';

// Assembles a full example program from snippet regions. The template is a
// plain string where "{{fileKey:region_name}}" placeholders are replaced by
// the named region from the matching file in `files`; the placeholder's own
// indentation is applied to every inserted line.
//
//   <GoExample
//     files={{authGo, tablesGo}}
//     template={`package main
//
//   {{tablesGo:table_imports}}
//
//   func main() {
//       {{authGo:auth_basic}}
//   }`}
//   />

const PLACEHOLDER = /^([\t ]*)\{\{([A-Za-z0-9_]+):([A-Za-z0-9_]+)\}\}\s*$/;

export default function GoExample({
  files,
  template,
  language = 'go',
  title,
}: {
  files: Record<string, string>;
  template: string;
  language?: string;
  title?: string;
}): React.ReactElement {
  const rendered = template
    .split('\n')
    .map((line) => {
      const m = line.match(PLACEHOLDER);
      if (!m) return line;
      const [, indent, fileKey, region] = m;
      const src = files[fileKey];
      if (src === undefined) throw new Error(`unknown snippet file key "${fileKey}"`);
      return extractRegion(src, region)
        .split('\n')
        .map((l) => (l.trim() === '' ? '' : indent + l))
        .join('\n');
    })
    .join('\n');
  return (
    <CodeBlock language={language} title={title}>
      {rendered}
    </CodeBlock>
  );
}
