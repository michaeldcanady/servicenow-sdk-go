#!/bin/bash

CURRENT_VERSION=$(cat VERSION | sed 's/^v//')
LAST_STABLE_VERSION=$(gh release list --limit 1 --json tagName --jq '.[0].tagName' | sed 's/^v//')

if [ $CURRENT_VERSION -gt $LATEST_RELEASE_COMMIT_ID ]; then
    echo true
else
    echo false
fi