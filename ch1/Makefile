.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./..

clean:
	go clean
	rm -f hello_world

build: vet
	go build

