.DEFAULT_GOAL := build

.PHONY: vet build

vet:
	go vet ./...

build: vet
	go build -o build/aoc-2020

clean:
	rm -rf build

