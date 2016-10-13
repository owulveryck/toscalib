#!/bin/bash

rm -f err.txt*
rm -f vet.txt

for pkg in $(glide nv); do
    if [[ $pkg != *"/doc/"* ]]; then
        errcheck $pkg >> err.txt
        go vet $pkg >> vet.txt 2>&1
    fi
done

# ignore generated files
sed -i.prev '/defer/d' err.txt

# remove when it exists
rm -f err.txt.prev

if [[ -s err.txt ]] || [[ -s vet.txt ]]
then
    cat err.txt
    cat vet.txt
    exit 1
fi
