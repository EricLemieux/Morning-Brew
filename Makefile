.PHONY: build
build:
	go build -o dist/morning-brew ./

.PHONY: test
test:
	go test -cover ./...

.PHONY: format
format:
	go fmt ./...