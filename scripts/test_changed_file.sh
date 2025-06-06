#!/bin/bash

LATEST_RELEASE_TAG=$(git describe --tags `git rev-list --tags --max-count=1`)

# Get the current state of the repo
current_state=$(git rev-parse HEAD)

echo "files: $(git diff --name-only "$LATEST_RELEASE_TAG" "$current_state" | grep -q '\.go$')"

if git diff --name-only "$LATEST_RELEASE_TAG" "$current_state" | grep -q '\.go$'; then
  exit 0
else
  exit 1
fi
