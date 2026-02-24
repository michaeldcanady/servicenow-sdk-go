#!/usr/bin/env bash
set -e

echo "Installing Podman client..."
apt-get update -y
apt-get install -y podman

echo "Podman remote client installed."
echo "Expecting Podman socket at: $SOCKET_PATH"
