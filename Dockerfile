# Dockerfile for http-to-grpc proxy server
# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

# copy dependencies
COPY ./ ./

# install dependencies
RUN go mod download

# compile server
RUN	mkdir -p bin
RUN	go build -o bin/server ./cmd/server/main.go

EXPOSE 8081

# run server binary
CMD ["./bin/server"]
