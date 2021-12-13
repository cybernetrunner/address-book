proto:
	buf generate

test:
	go test ./... -race

env:
	export $(grep -v '^#' .env | xargs)

depend:
	go mod tidy -go=1.16
	go mod vendor

run:
	go fmt ./...
	golangci-lint run ./...
	echo "DEBUG: RUN SERVER"
	go run cmd/server/main.go

build:
	go fmt ./...
	golangci-lint run ./...
	mkdir -p bin
	go build -o bin/server cmd/server/main.go

clear:
	rm -r gen
	rm -r bin

deploy:
	minikube start --vm-driver=hyperkit
	minikube status
	kubectl apply -f k8s --validate=false
	kubectl top node
	minikube dashboard

migrateup:
	migrate -path internal/database/migrations -database "postgresql://${DB_USER}:${DB_NAME}@${DB_HOST}:${DB_PORT}/go_sample?sslmode=${DB_SSLMODE}" -verbose up

migratedown:
	migrate -path internal/database/migrations -database "postgresql://${DB_USER}@${DB_NAME}:${DB_PORT}/go_sample?sslmode=${DB_SSLMODE}" -verbose down
