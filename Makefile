MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc
.DEFAULT_GOAL := test

.PHONY: clean test cover lint format

PROJECT = toscalib

IMPORT_PATH := github.com/CiscoCloud/${PROJECT}

ROOT ?= /${PROJECT}
DEV_IMAGE := ${PROJECT}_dev

DOCKERRUN := docker run --rm \
	-v ${ROOT}/vendor:/go/src \
	-v ${ROOT}:/${PROJECT}/src/${IMPORT_PATH} \
	-w /${PROJECT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}

DOCKERNOVENDOR := docker run --rm -i \
	-e LDFLAGS="${LDFLAGS}" \
	-v ${ROOT}:/${PROJECT}/src/${IMPORT_PATH} \
	-w /${PROJECT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}

clean:
	@rm -rf cover *.txt

# ----------------------------------------------
# docker build

# builds the builder container
build/image_build:
	docker build -t ${DEV_IMAGE} -f Dockerfile.dev .

# top-level target for vendoring our packages: glide install requires
# being in the package directory so we have to run this for each package
vendor: build/image_build
	${DOCKERRUN} glide install

# fetch a dependency via go get, vendor it, and then save into the parent
# package's glide.yml
# usage DEP=github.com/owner/package make add-dep
add-dep: build/image_build
	${DOCKERNOVENDOR} bash -c "DEP=$(DEP) ./scripts/add_dep.sh"

# ----------------------------------------------
# develop and test

format: vendor
	${DOCKERNOVENDOR} bash ./scripts/fmt.sh

check: format
	${DOCKERRUN} bash ./scripts/check.sh

# default task
test: check
	${DOCKERRUN} bash ./scripts/test.sh

# run unit tests and write out test coverage
cover: check
	@rm -rf cover/
	@mkdir -p cover
	${DOCKERRUN} bash ./scripts/cover.sh

# ------ Generator
generate: build/image_build NormativeTypes/*
	${DOCKERRUN} go-bindata -pkg=toscalib -prefix=NormativeTypes/ -o normative_definitions.go NormativeTypes/
