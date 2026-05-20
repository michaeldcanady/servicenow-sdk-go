#!/usr/bin/env bash
set -euo pipefail

echo "Installing Antigravity CLI (headless)..."

export DEBIAN_FRONTEND=noninteractive

# Minimal dependencies
apt-get update -y
apt-get install -y --no-install-recommends curl ca-certificates tar
rm -rf /var/lib/apt/lists/*

# Install directory for the CLI
INSTALL_DIR="/usr/local/bin"

# Run the official installer with explicit install directory
curl -fsSL https://antigravity.google/cli/install.sh | bash -s -- --dir "$INSTALL_DIR"

# The installer places the binary as "agy"
CLI_BIN="$INSTALL_DIR/agy"

if [ ! -x "$CLI_BIN" ]; then
    echo "ERROR: Antigravity CLI binary not found at $CLI_BIN" >&2
    exit 1
fi

echo "Antigravity CLI installed successfully at $CLI_BIN"
"$CLI_BIN" --version || true
