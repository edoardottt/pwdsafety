PROJECT_NAME := "pwdsafety"
PKG := "github.com/edoardottt/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

fmt:
	@gofmt -s ./*;
	@echo "Done."

remod:
	@rm -rf go.*
	@go mod init ${PKG}
	@cd cmd && go get
	@echo "Done."

update:
	@cd cmd && go get -u;
	@cd cmd && go mod tidy -v;
	@echo "Done."

linux:
	@cd cmd && go build -o ./pwdsafety
	@sudo mv ./cmd/pwdsafety /usr/bin/
	@echo "Done."

unlinux:
	@sudo rm -rf /usr/bin/pwdsafety
	@echo "Done."

test:
	@go test -v -race ./...
	@echo "Done."

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}
