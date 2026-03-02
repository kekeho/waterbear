SHELL := /bin/bash

GOPATH ?= $(CURDIR)
GOBIN ?= $(CURDIR)/bin
GO111MODULE ?= off
export GOPATH
export GOBIN
export GO111MODULE

.PHONY: go install build sync-repo

go:
	@$(MAKE) sync-repo REPO=https://github.com/grpc/grpc-go.git DEST="$(GOPATH)/src/google.golang.org/grpc"
	@$(MAKE) sync-repo REPO=https://github.com/golang/net.git DEST="$(GOPATH)/src/golang.org/x/net"
	@$(MAKE) sync-repo REPO=https://github.com/golang/text.git DEST="$(GOPATH)/src/golang.org/x/text"
	@$(MAKE) sync-repo REPO=https://github.com/golang/crypto.git DEST="$(GOPATH)/src/golang.org/x/crypto"
	@$(MAKE) sync-repo REPO=https://github.com/golang/sys.git DEST="$(GOPATH)/src/golang.org/x/sys"
	@$(MAKE) sync-repo REPO=https://github.com/googleapis/go-genproto.git DEST="$(GOPATH)/src/google.golang.org/genproto"
	@$(MAKE) sync-repo REPO=https://github.com/protocolbuffers/protobuf-go.git DEST="$(GOPATH)/src/google.golang.org/protobuf"

install:
	@$(MAKE) sync-repo REPO=https://github.com/klauspost/reedsolomon.git DEST="$(GOPATH)/src/github.com/klauspost/reedsolomon"
	@$(MAKE) sync-repo REPO=https://github.com/klauspost/cpuid.git DEST="$(GOPATH)/src/github.com/klauspost/cpuid"
	@$(MAKE) sync-repo REPO=https://github.com/cbergoon/merkletree.git DEST="$(GOPATH)/src/github.com/cbergoon/merkletree"
	@$(MAKE) sync-repo REPO=https://github.com/golang/protobuf.git DEST="$(GOPATH)/src/github.com/golang/protobuf"

build:
	go install src/main/server.go
	go install src/main/client.go
	go install src/main/keygen.go

sync-repo:
	@if [ -z "$(REPO)" ] || [ -z "$(DEST)" ]; then \
		echo "sync-repo requires REPO and DEST"; \
		exit 1; \
	fi
	@if [ -d "$(DEST)/.git" ]; then \
		echo "Updating $(DEST)"; \
		git -C "$(DEST)" pull --ff-only; \
	else \
		echo "Cloning $(REPO) -> $(DEST)"; \
		mkdir -p "$(dir $(DEST))"; \
		git clone --depth 1 "$(REPO)" "$(DEST)"; \
	fi
	
