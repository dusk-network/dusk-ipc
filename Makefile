PROJECT_NAME := "dusk-ipc"
PKG := "github.com/dusk-network/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
#TEST_FLAGS := "-count=1"
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
.PHONY: all go clean gen regen-dir help

go: gen-go-grpc generate ## Run all Go generators in this package
	@echo "generated packages"
all: go ## (re)Generate all the things
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
gen-go-grpc:
	@protoc -I./ipc/ ./ipc/*.proto --go-grpc_out=paths=source_relative:./ipc/ --go_out=paths=source_relative:./ipc/

generate:
	@go generate ./...
