#!/usr/bin/env bash
set -e

echo "Updating package lists..."
apt-get update

echo "Installing Podman client..."
apt-get install -y --no-install-recommends podman
rm -rf /var/lib/apt/lists/*

echo "Podman remote client installed."
