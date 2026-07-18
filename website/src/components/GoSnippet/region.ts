// Extracts a named region from a Go snippet source file. Regions are
// delimited by "// [START name]" / "// [END name]" marker comments, the
// same markers the old mkdocs include-markdown setup consumed.

export function extractRegion(src: string, region: string): string {
  const lines = src.split('\n');
  const start = lines.findIndex((l) => l.includes(`[START ${region}]`));
  const end = lines.findIndex((l) => l.includes(`[END ${region}]`));
  if (start === -1 || end === -1 || end <= start) {
    throw new Error(`snippet region "${region}" not found`);
  }
  return dedent(lines.slice(start + 1, end));
}

function dedent(lines: string[]): string {
  const indents = lines
    .filter((l) => l.trim() !== '')
    .map((l) => l.match(/^[\t ]*/)![0].length);
  const min = indents.length ? Math.min(...indents) : 0;
  return lines
    .map((l) => l.slice(min))
    .join('\n')
    .replace(/^\n+|\n+$/g, '');
}
