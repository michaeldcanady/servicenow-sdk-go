#!/bin/bash

CURRENT_VERSION_STRING=$(cat VERSION | sed 's/^v//')
LAST_STABLE_VERSION_STRING="$(git tag | grep -v -- '-' | sort -V | tail -n 1 | sed 's/^v//')"

if [[ $CURRENT_VERSION_STRING == $LAST_STABLE_VERSION_STRING ]]
then
    exit 1
fi
IFS=.
ver1=($CURRENT_VERSION_STRING)
ver2=($LAST_STABLE_VERSION_STRING)
# fill empty fields in ver1 with zeros
for ((i=${#ver1[@]}; i<${#ver2[@]}; i++))
do
    ver1[i]=0
done
for ((i=0; i<${#ver1[@]}; i++))
do
    if ((10#${ver1[i]:=0} > 10#${ver2[i]:=0}))
    then
        exit 0
    fi
    if ((10#${ver1[i]} < 10#${ver2[i]}))
    then
        exit 1
    fi
done
exit 1
