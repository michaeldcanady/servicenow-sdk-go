#!/bin/bash

LATEST_RELEASE_TAG=$(gh release list --limit 1 --json tagName --jq '.[0].tagName')
LATEST_RELEASE_COMMIT_ID=$(git rev-parse "${LATEST_RELEASE_TAG}")
CURRENT_COMMIT_ID=$(git rev-parse HEAD)

if ["$CURRENT_COMMIT_ID" != "$LATEST_RELEASE_COMMIT_ID"]; then
    echo true
else
    echo false
fi