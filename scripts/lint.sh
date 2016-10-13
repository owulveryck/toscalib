#!/bin/bash

rm -f lint.txt*

for pkg in $(glide nv); do
    golint $pkg >> lint.txt
done

# ignore generated files
sed -i.prev '/normative_definitions.go/d' lint.txt

# remove when it exists
rm -f lint.txt.prev

if [[ -s lint.txt ]]
then
    cat lint.txt
    exit 1
fi
