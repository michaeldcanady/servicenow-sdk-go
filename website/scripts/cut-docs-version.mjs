// Cuts a docs version snapshot and registers it in the site config.
// Usage: node scripts/cut-docs-version.mjs <X.Y>
//
// Used by .github/workflows/docs-version.yml on stable release tags, and
// safe to run manually as a fallback (see "Docs site versions" in
// docs/contributing/release-branches.md).

import {execSync} from 'node:child_process';
import fs from 'node:fs';
import path from 'node:path';
import {fileURLToPath} from 'node:url';

const websiteDir = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '..',
);

const version = process.argv[2];
if (!version || !/^\d+\.\d+$/.test(version)) {
  console.error('Usage: node scripts/cut-docs-version.mjs <X.Y>');
  process.exit(1);
}

const versionedDir = path.join(websiteDir, 'versioned_docs', `version-${version}`);
if (fs.existsSync(versionedDir)) {
  console.error(`Docs version ${version} already exists; nothing to do.`);
  process.exit(1);
}

// Validate the config anchor before mutating anything: if the check ran
// after `docs:version`, a config drift would leave a half-registered
// version behind (created dirs, versions.json entry) and a dirty tree.
const configPath = path.join(websiteDir, 'docusaurus.config.ts');
const configLines = fs.readFileSync(configPath, 'utf8').split('\n');
const anchorIdx = configLines.findIndex((line) =>
  line.includes("current: {label: 'main'"),
);
if (anchorIdx === -1) {
  console.error(
    "Could not find the versions anchor (current: {label: 'main'...}) in docusaurus.config.ts; add the version entry manually.",
  );
  process.exit(1);
}

execSync(`npm run --silent docusaurus docs:version ${version}`, {
  cwd: websiteDir,
  stdio: 'inherit',
});

// Register the version in the config so the dropdown label is "vX.Y" and
// the released line carries no "Unreleased" banner.
configLines.splice(
  anchorIdx + 1,
  0,
  `            '${version}': {label: 'v${version}', banner: 'none'},`,
);
fs.writeFileSync(configPath, configLines.join('\n'));

console.log(`Docs version ${version} cut and registered.`);
