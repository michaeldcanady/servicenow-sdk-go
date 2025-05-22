#!/bin/bash

LATEST_RELEASE_TAG=$(gh release list --limit 1 --json tagName --jq '.[0].tagName')

# Get the current state of the repo
current_state=$(git rev-parse HEAD)

echo $(git diff --name-only $LATEST_RELEASE_TAG $current_state | grep '\.go$' || true)
