
tests: cleaning
	go test ./... -race

env:
	export $(grep -v '^#' .env | xargs)

depend:
	go mod tidy -go=1.16
	go mod vendor

protogen:
	buf generate

run: cleaning
	echo "DEBUG: RUN SERVER"
	go run cmd/server/main.go

build: cleaning
	mkdir -p bin
	go build -o bin/server cmd/server/main.go

clear:
	rm -r gen
	rm -r bin

cleaning:
	go fmt ./...
	golangci-lint run ./...

deploy:
	minikube start
	minikube status
	kubectl apply -f k8s --validate=false
	kubectl get services
	minikube dashboard

compose_up:
	echo "DEBUG: RUN DOCKER-COMPOSE (DB+SERVER)"
	docker-compose --env-file .env build
	docker-compose --env-file .env up

compose_down:
	docker-compose down

migrate_up: env
	migrate -path internal/database/migrations -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}" -verbose up

migrate_down: env
	migrate -path internal/database/migrations -database "postgresql://${DB_USER}@${DB_PASSWORD}:${DB_PORT}/go_sample?sslmode=${DB_SSLMODE}" -verbose down
