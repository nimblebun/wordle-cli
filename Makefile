default: build

build:
	mkdir -p dist
	go build -o dist

.PHONY: build
