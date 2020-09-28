BINARY := dotenvtok8s
OSES := darwin linux
ARCHS := amd64
OSANDARCH := $(foreach o,$(OSES),$(foreach a,$(ARCHS),$(o:%=$(BINARY)-%-$(a))))
BUILD ?= $(shell (git tag --points-at HEAD | tr '\n' ' ';git rev-parse HEAD | tr -d '\n'))
LDFLAGS = -ldflags '-X "main.version=$(BUILD)"'

.PHONY: build-all
build-all: $(OSANDARCH:%=build/%)

.PHONY: package
package: $(OSANDARCH:%=build/%.tar.gz)

build/%.tar.gz: build/%
	rm -f build/$(BINARY) && cp -v build/$(<:build/%=%) build/$(BINARY)
	tar -cvzf $@ -C build $(BINARY)

build/%-darwin-amd64: cmd/%.go build
	GOOS=darwin ARCH=amd64 go build -o $@ $(LDFLAGS) $<

build/%-linux-amd64: cmd/%.go build
	GOOS=linux ARCH=amd64 go build -o $@ $(LDFLAGS) $<

clean:
	rm -rf ./build/*

build:
	mkdir -p ./build

.PHONY: test
test:
	go test -v ./...
