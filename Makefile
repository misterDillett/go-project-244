.PHONY: build run test lint clean fmt clean-cache

build:
	go build -o bin/gendiff cmd/gendiff/main.go

run:
	go run cmd/gendiff/main.go

test:
	go test -v ./...

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin/

clean-cache:
	go clean -testcache
	rm -f go.work go.work.sum 2>/dev/null || true