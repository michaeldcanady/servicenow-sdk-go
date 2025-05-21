#!/bin/bash

LATEST_RELEASE_TAG=$(gh release list --limit 1 --json tagName --jq '.[0].tagName')
echo $(git diff --name-only $LATEST_RELEASE_TAG HEAD | grep '\.go$' || true)
