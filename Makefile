MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc
.DEFAULT_GOAL := test

export PROJECT = toscalib

IMPORT_PATH := github.com/CiscoCloud/${PROJECT}

# Windows environment?
CYG_CHECK := $(shell hash cygpath 2>/dev/null && echo 1)
ifeq ($(CYG_CHECK),1)
	VBOX_CHECK := $(shell hash VBoxManage 2>/dev/null && echo 1)

	# Docker Toolbox (pre-Windows 10)
	ifeq ($(VBOX_CHECK),1)
		ROOT := /${PROJECT}
	else
		# Docker Windows
		ROOT := $(shell cygpath -m -a "$(shell pwd)")
	endif
else
	# all non-windows environments
	ROOT := $(shell pwd)
endif

INSTALL_VENDOR := $(shell [ ! -d vendor ] && echo 1)
FORCE_VENDOR_INSTALL ?=
ifneq ($(strip $(FORCE_VENDOR_INSTALL)),)
	INSTALL_VENDOR := 1
endif

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
	@rm -rf cover

veryclean: clean
	@rm -rf .glide vendor

# ----------------------------------------------
# build

# builds the builder container
.PHONY: build/image_build
build/image_build:
	@echo "Building dev container"
	@docker build --quiet -t ${DEV_IMAGE} -f Dockerfile.dev .

# top-level target for vendoring our packages: glide install requires
# being in the package directory so we have to run this for each package
.PHONY: vendor
vendor: build/image_build
ifeq ($(INSTALL_VENDOR),1)
	${DOCKERRUN} glide install --skip-test
endif

# fetch a dependency via go get, vendor it, and then save into the parent
# package's glide.yml
# usage: DEP=github.com/owner/package make add-dep
.PHONY: add-dep
add-dep: build/image_build
ifeq ($(strip $(DEP)),)
	$(error "No dependency provided. Expected: DEP=<go import path>")
endif
	${DOCKERNOVENDOR} glide get --skip-test ${DEP}

# ----------------------------------------------
# develop and test

.PHONY: format
format: vendor
	${DOCKERNOVENDOR} bash ./scripts/fmt.sh

.PHONY: check
check: format
	${DOCKERNOVENDOR} bash ./scripts/check.sh

# default task
.PHONY: test
test: check
	${DOCKERRUN} bash ./scripts/test.sh

# run unit tests and write out test coverage
.PHONY: cover
cover: check
	@rm -rf cover/
	@mkdir -p cover
	${DOCKERRUN} bash ./scripts/cover.sh

# ------ Generator
.PHONY: generate
generate: build/image_build NormativeTypes/*
	${DOCKERRUN} go-bindata -pkg=toscalib -prefix=NormativeTypes/ -o normative_definitions.go NormativeTypes/

# ------ Minishift / Docker Machine Helpers
.PHONY: setup
setup:
	@bash ./scripts/setup.sh
