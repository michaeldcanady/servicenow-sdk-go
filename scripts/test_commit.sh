#!/bin/bash

LATEST_RELEASE_TAG=$(git describe --tags `git rev-list --tags --max-count=1`)
LATEST_RELEASE_COMMIT_ID=$(git rev-parse "${LATEST_RELEASE_TAG}")
CURRENT_COMMIT_ID=$(git rev-parse HEAD)

if [ "$CURRENT_COMMIT_ID" != "$LATEST_RELEASE_COMMIT_ID" ]; then
    exit 0
else
    exit 1
fi