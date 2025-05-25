#!/bin/bash

git fetch --quiet --tags origin

LATEST_RELEASE_TAG=$(git describe --tags `git rev-list --tags --max-count=1`)

# Get the current state of the repo
current_state=$(git rev-parse HEAD)

echo $(git diff --name-only $LATEST_RELEASE_TAG $current_state | grep '\.go$' || true)
