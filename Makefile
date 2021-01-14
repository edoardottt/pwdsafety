PROJECT_NAME := "pwdsafety"
PKG := "github.com/edoardottt/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

fmt:
	@gofmt -s ./*; \
	echo "Done."

remod:
	rm -rf go.*
	go mod init ${PKG}
	go get
	@echo "Done."

update:
	@go get -u; \
	go mod tidy -v; \
	echo "Done."

linux:
	@go build -o ./pwdsafety
	mv ./pwdsafety /usr/bin/
	mkdir -p /usr/bin/pwds/
	cp -r pwds/*.txt /usr/bin/pwds/
	@echo "Done."

unlinux:
	rm -rf /usr/bin/pwdsafety
	rm -rf /usr/bin/pwds
	@echo "Done."

test:
	@go test -v -race ./... ; \
	echo "Done."

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

build: dep ## Build the binary file
	@go build -i -o build/main $(PKG)
