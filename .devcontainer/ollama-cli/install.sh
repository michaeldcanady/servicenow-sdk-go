#!/usr/bin/env bash
set -e

echo "Installing dependencies..."
apt-get update
apt-get install -y --no-install-recommends zstd curl ca-certificates
rm -rf /var/lib/apt/lists/*

echo "Installing Ollama CLI..."
# Use the official installer but we only want the binary
curl -fsSL https://ollama.com/install.sh | sh

# The installer might try to start a service, but we don't need it in this container
# since the server is in the 'ollama' service.
systemctl disable ollama || true

# Remove OLLAMA_HOST from /etc/environment if the installer added it, 
# so the feature's containerEnv can correctly point to the server service.
sed -i '/OLLAMA_HOST/d' /etc/environment

echo "Ollama CLI installed successfully."
