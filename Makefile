create:
	buf generate

clear:
	rm gen/proto/*.go
	rm bin/*

depend:
	go mod tidy -go=1.16
	go mod vendor

test:
	go test
	go test -race

build:
	go fmt ./...
	golangci-lint run ./...
	go build -o bin/address-book cmd/address-book/main.go

run:
	go fmt ./...
	golangci-lint run ./...
	go run cmd/address-book/main.go
