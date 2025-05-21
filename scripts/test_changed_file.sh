#!/bin/bash

LATEST_RELEASE_TAG=$(gh release list --limit 1 --json tagName --jq '.[0].tagName')
return $(git diff --name-only $LATEST_RELEASE_TAG HEAD | grep '\.go$' || true)
