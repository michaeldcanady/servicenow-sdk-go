#!/usr/bin/env bash
set -e

echo "Updating package lists..."
apt-get update

echo "Installing xdg-utils..."
apt-get install -y --no-install-recommends xdg-utils
rm -rf /var/lib/apt/lists/*
