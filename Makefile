SHELL:=/bin/sh
.PHONY: all build clean docker format test vet

export GO111MODULE=on
export GOPROXY=https://goproxy.io

pkgs = $(shell go list ./...)

MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR  := $(dir $(MKFILE_PATH))

SOURCE = $(shell find ${MKFILE_DIR} -path "${MKFILE_DIR}vendor" -prune -o -type f -name "*.go" -print)
TARGET = ${MKFILE_DIR}/prom-elastic-alert

all: ${TARGET}

${TARGET}: ${SOURCE}
	@echo ">> building code"
	go mod tidy
	go mod vendor
	go build -o ${TARGET}

build: all

clean:
	@echo ">> cleaning build"
	rm  ${TARGET}

format:
	@echo ">> formatting code"
	go fmt $(pkgs)

test:
	@echo ">> running short tests"
	go test -v $(pkgs)

vet:
	@echo ">> vetting code"
	go vet $(pkgs)
