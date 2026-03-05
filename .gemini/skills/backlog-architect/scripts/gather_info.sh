#!/bin/bash

echo "---"
echo "OS: $(uname -a)"
echo "Go version: $(go version)"
if [ -f VERSION ]; then
    echo "Module version: $(cat VERSION)"
else
    echo "Module version: Unknown (VERSION file not found)"
fi
echo "---"