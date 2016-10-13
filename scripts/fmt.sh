#!/bin/bash

find . \( -path ./vendor -o -path ./.glide \) -prune -o -name "*.go" -exec goimports -w {} \;
