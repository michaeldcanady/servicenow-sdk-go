#!/bin/bash
# scripts/replace_unreleased.sh
# Replaces {unreleased} with a given tag in all .go files.

TAG=$1

if [ -z "$TAG" ]; then
  echo "Usage: $0 <tag>"
  exit 1
fi

echo "Replacing {unreleased} with $TAG in all .go files..."
find . -name "*.go" -type f -exec sed -i "s/{unreleased}/${TAG}/g" {} \;
