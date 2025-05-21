#!/bin/bash

unreleased_version=$(cat VERSION)
current_date=$(date +'%Y%m%d')
echo "$unreleased_version-preview$current_date"