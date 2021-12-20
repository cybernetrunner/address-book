
tests: cleaning
	go test ./... -race

depend:
	go mod tidy -go=1.16
	go mod vendor

protogen:
	buf generate

run:
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

migrate_up:
	migrate -path internal/database/migrations -database "postgresql://gorm:gorm12345@10.96.0.3:5432/gorm?sslmode=disable" -verbose up

migrate_down:
	migrate -path internal/database/migrations -database "postgresql://gorm@gorm12345:5432/gorm?sslmode=disable" -verbose down
