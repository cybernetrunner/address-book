create:
	mkdir -p google/api
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto

	#buf generate

	protoc -I . \
    	--go_out ./gen --go_opt paths=source_relative \
    	--go-grpc_out ./gen --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./gen --grpc-gateway_opt paths=source_relative \
    	./proto/api.proto

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
