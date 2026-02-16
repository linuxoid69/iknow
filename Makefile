APP=iknow
VERSION=0.4.0
REVISION=-0
GO_LDFLAGS=-ldflags="-s -w"

.PHONY: all
all:
	@echo 'DEFAULT:      '
	@echo '   make build   - build application'
	@echo '   make tools   - install tools to local ./bin '
	@echo '   make tidy    - go mod tidy'
	@echo '   make lint    - run linter'

.PHONY: build
build:
	@go build -trimpath -ldflags=$(GO_LDFLAGS)

.PHONY: tools
tools:
	@GOBIN=$PWD/bin go install tool

.PHONY: tidy
tidy:
	@go mod tidy
