#!/bin/bash

go test -v $(glide nv) -bench .
