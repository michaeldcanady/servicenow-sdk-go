#!/bin/bash

version_to_number() {
    echo "$1" | awk -F. '{ printf("%d%d%d\n", $1,$2,$3) }'
}

echo "$(git tag | grep -v -- '-' | sort -V | tail -n 1)"

CURRENT_VERSION_STRING=$(cat VERSION | sed 's/^v//')
LAST_STABLE_VERSION_STRING=$(gh release list --limit 1 --json tagName --jq '.[0].tagName' | sed 's/^v//')

CURRENT_VERSION=$(version_to_number $CURRENT_VERSION_STRING)
LAST_STABLE_VERSION=$(version_to_number $LAST_STABLE_VERSION_STRING)

echo "$CURRENT_VERSION"
echo "$LAST_STABLE_VERSION"

if [ $CURRENT_VERSION -gt $LAST_STABLE_VERSION ]; then
    echo true
else
    echo false
fi