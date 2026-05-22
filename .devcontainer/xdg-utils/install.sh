#!/usr/bin/env bash
set -e

# Use the option from devcontainer-feature.json
# Options are passed as environment variables prefixed with the option name in uppercase.
CONTAINER_PIPE="${CONTAINERBROWSERPIPEPATH:-/tmp/hostbrowserpipe}"

echo "Updating package lists..."
apt-get update

echo "Installing xdg-utils..."
apt-get install -y --no-install-recommends xdg-utils

# Create a wrapper for xdg-open to use the host browser pipe if available
echo "Creating xdg-open wrapper targeting ${CONTAINER_PIPE}..."
cat << EOF > /usr/local/bin/xdg-open
#!/usr/bin/env bash
# If the host browser pipe exists, send the URL to it.
# Otherwise, fall back to the real xdg-open.
if [ -p "${CONTAINER_PIPE}" ]; then
    echo "\$1" > "${CONTAINER_PIPE}"
    echo "URL sent to host browser pipe: \$1"
else
    if [ -f /usr/bin/xdg-open ]; then
        /usr/bin/xdg-open "\$@"
    else
        echo "Error: /usr/bin/xdg-open not found and ${CONTAINER_PIPE} is not a pipe."
        exit 1
    fi
fi
EOF
chmod +x /usr/local/bin/xdg-open

rm -rf /var/lib/apt/lists/*
