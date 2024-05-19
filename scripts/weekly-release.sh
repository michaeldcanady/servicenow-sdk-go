#!/bin/bash

# Function to split a version into major, minor, and patch versions
split_version() {
    version=$1
    echo $(echo $version | cut -d. -f1)
    echo $(echo $version | cut -d. -f2)
    echo $(echo $version | cut -d. -f3)
}

# Get the latest release
latest_release=$(git describe --tags `git rev-list --tags --max-count=1`)

# Get the latest STABLE release
latest_stable_release=$(git tag | grep -v -- '-preview' | sort -V | tail -n 1)

# Get the current state of the repo
current_state=$(git rev-parse HEAD)

# Compare the last release to the current state of the repo
file_changes=$(git diff --name-only $latest_release $current_state | grep -v "internal" | grep ".go$")

# Check if ANY file changes have occurred
if [[ $file_changes == "" ]]; then 
    echo "::set-output name=changed::false"
    exit 0 
fi

# Get the latest STABLE version
latest_stable_version=$(echo $latest_stable_release | sed 's/v//')
{ 
    read latest_stable_major_version
    read latest_stable_minor_version
    read latest_stable_patch_version
} <<< $(split_version $latest_stable_version)

# Get the current version
if [[ ! -f ./VERSION ]]; then
    echo "Error: VERSION file not found"
    exit 1
fi

# Get the Current version
UNRELEASED_VERSION=$(cat ./VERSION)
{
    read current_major_version
    read current_minor_version
    read current_patch_version
} <<< $(split_version $current_version)

# Check that current version is greater than latest stable in some way
if [[ $latest_stable_major_version < $current_major_version ]]; then
    echo "::set-output name=changed::true"
    exit 0 
fi

if [[ $latest_stable_minor_version < $current_minor_version ]]; then
    echo "::set-output name=changed::true"
    exit 0 
fi

if [[ $latest_stable_patch_version < $current_patch_version ]]; then
    echo "::set-output name=changed::true"
else
    echo "::set-output name=changed::false"
fi
