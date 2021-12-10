create:
	buf generate

clear:
	rm -r gen
	rm -r bin

depend:
	go mod tidy -go=1.16
	go mod vendor

test:
	go test ./... -race

build:
	go fmt ./...
	golangci-lint run ./...
	mkdir -p bin
	go build -o bin/server cmd/server/main.go

run:
	go fmt ./...
	golangci-lint run ./...
	go run cmd/server/main.go
