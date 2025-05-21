#!/bin/bash

CURRENT_VERSION=$(cat VERSION | sed 's/^v//')
LAST_STABLE_VERSION=$(gh release list --limit 1 --json tagName --jq '.[0].tagName' | sed 's/^v//')

return $CURRENT_VERSION > $LAST_STABLE_VERSION
