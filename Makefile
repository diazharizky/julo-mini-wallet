###################
### DEVELOPMENT ###
###################

.PHONY: tidy run migrate-up migrate-down

tidy:
	go mod tidy

run:
	go run cmd/main.go

migrate-up:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path /migrations -database postgres://julo:julo@localhost:5432/julo?sslmode=disable -verbose up

migrate-down:
	docker run -v ./migrations:/migrations --network host migrate/migrate -path /migrations -database postgres://julo:julo@localhost:5432/julo?sslmode=disable -verbose down -all

#############
### BUILD ###
#############

.PHONY: compile build

compile:
	rm -rf bin/app && \
	go build -v -o bin/app ./cmd

build:
	docker build -t julo-mini-wallet:latest .
