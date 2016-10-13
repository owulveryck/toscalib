#!/bin/bash
if [ -z "$DEP" ]; then
  echo "No dependency provided. Expected: DEP=<go import path>"
  exit 1
fi
glide get ${DEP}
