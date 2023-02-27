GO = go

.PHONY: upgrade build release test

upgrade: build test release

build:
	./script/builder.sh

release:
	./script/release.sh

test:
	$(GO) test -v -cover -race -coverprofile=cover.out .
	$(GO) tool cover -html=cover.out
