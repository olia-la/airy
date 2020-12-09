#!/bin/bash
set -eo pipefail
IFS=$'\n\t'

echo "Running Bazel lint"
bazel run @com_github_airyhq_bazel_tools//code-format:check_buildifier
echo
echo "Running Prettier and Java tests"
bazel test --test_tag_filters=lint //...

go get -u golang.org/x/lint/golint

if [[ $(golint ./...) ]]; then
    exit 0
else
    exit 1
fi