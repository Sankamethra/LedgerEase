PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell git describe --tags)
COMMIT := $(shell git log -1 --format='%H')
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=invoice \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=invoice \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'
GOPROXY=https://proxy.golang.org GO111MODULE=on

all: install

install: go.sum
	@echo "--> Installing invoice"
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/invoiced

build: go.sum
	@echo "--> Building invoice"
	@go build -mod=readonly $(BUILD_FLAGS) ./cmd/invoiced

build-linux: go.sum
	@echo "--> Building Linux invoice"
	GOOS=linux GOARCH=amd64 go build -mod=readonly $(BUILD_FLAGS) ./cmd/invoiced

build-darwin: go.sum
	@echo "--> Building macOS invoice"
	go build -mod=readonly $(BUILD_FLAGS) ./cmd/invoiced

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
