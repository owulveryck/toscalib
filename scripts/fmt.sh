#!/bin/bash

find . \( -path ./vendor -o -path ./.glide \) -prune -o -name "*.go" -exec goimports -w {} \;

if ! git diff-index --quiet HEAD; then
    echo "goimports modified code; requires attention!"
    if [[ "${CI}" == "true" ]]; then
        exit 1
    fi
fi
