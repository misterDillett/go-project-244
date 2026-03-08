.PHONY: build run test lint clean fmt

build:
	go build -o bin/gendiff cmd/gendiff/main.go

run:
	go run cmd/gendiff/main.go

test:
	go test -v .

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin/