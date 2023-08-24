.PHONY: tidy run migrate-up migrate-down

tidy:
	go mod tidy

run:
	go run cmd/main.go

migrate-up:
	migrate -database postgres://julo:julo@localhost:5432/julo?sslmode=disable -path ./migrations -verbose up

migrate-down:
	migrate -database postgres://julo:julo@localhost:5432/julo?sslmode=disable -path ./migrations -verbose down
