#!/bin/bash

echo "$(git tag | grep -v -- '-' | sort -V | tail -n 1)"

CURRENT_VERSION_STRING=$(cat VERSION | sed 's/^v//')
LAST_STABLE_VERSION_STRING=$(gh release list --limit 1 --json tagName --jq '.[0].tagName' | sed 's/^v//')

echo "\n$CURRENT_VERSION_STRING\n"
echo "\n$LAST_STABLE_VERSION_STRING\n"

if [[ $CURRENT_VERSION_STRING == $LAST_STABLE_VERSION_STRING ]]
then
    echo false
fi
local IFS=.
local i ver1=($CURRENT_VERSION_STRING) ver2=($LAST_STABLE_VERSION_STRING)
# fill empty fields in ver1 with zeros
for ((i=${#ver1[@]}; i<${#ver2[@]}; i++))
do
    ver1[i]=0
done
for ((i=0; i<${#ver1[@]}; i++))
do
    if ((10#${ver1[i]:=0} > 10#${ver2[i]:=0}))
    then
        echo true
    fi
    if ((10#${ver1[i]} < 10#${ver2[i]}))
    then
        echo false
    fi
done
echo false

