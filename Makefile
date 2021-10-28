#
# simple makefile to build and release k2env
#

PWD                       := $(shell pwd)
PREFIX                    ?= $(GOPATH)
BINDIR                    ?= $(PREFIX)/bin
GO                        := GO111MODULE=on go
GOOS                      ?= $(shell go version | cut -d' ' -f4 | cut -d'/' -f1)
GOARCH                    ?= $(shell go version | cut -d' ' -f4 | cut -d'/' -f2)


build:
	$(GO) build -o k2env  cmd/k2env/main.go