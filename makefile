GO = go

.PHONY: upgrade build test

upgrade: build test

build:
	./script/builder.sh

test:
	$(GO) test -v -cover -race -coverprofile=cover.out .
	$(GO) tool cover -html=cover.out
