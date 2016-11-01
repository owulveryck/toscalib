#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Excludes:
#   - normative_definitions.go is generated code so static analysis will always have some issue
#   - when using defer there is no way to check to returned value so ignore
gometalinter \
    --exclude='normative_definitions\.go:.*$' \
    --exclude='error return value not checked.*(Close|Log|Print).*\(errcheck\)$' \
    --disable=aligncheck \
    --disable=dupl \
    --disable=gotype \
    --enable=unused \
    --cyclo-over=20 \
    --tests \
    --deadline=20s
