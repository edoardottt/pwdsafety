PROJECT_NAME := "pwdsafety"
PKG := "github.com/edoardottt/$(PROJECT_NAME)"

fmt:
	@gofmt -s ./*
	@echo "Done."

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
	cp engWordsList.txt /usr/bin/
	@echo "Done."

unlinux:
	rm -rf /usr/bin/pwdsafety
	rm -rf /usr/bin/pwds
	rm -rf /usr/bin/engWordsList.txt
	@echo "Done."

test:
	@go test -v -race ./... ; \
	echo "Done."
