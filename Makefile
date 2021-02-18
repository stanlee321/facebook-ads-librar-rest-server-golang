PROJECT_NAME = facebook-ads
BIN_DIR = ./bin
BIN_FILE = $(PROJECT_NAME)
CMD_DIR = ./cmd

# Get version constant
VERSION := 1.0.0
BUILD := $(shell git rev-parse HEAD)

# Use linker flags to provide version/build settings to the binary
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION) -X=main.build=$(BUILD)"

# Some ENV VARS

export DATABASE_DEV_URL = postgresql://root:secret@localhost:5435/stats_db?sslmode=disable


gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pkg/api/v1 --grpc-gateway_out=:pkg/api/v1 --openapiv2_out=:openapiv2

postgres:
	docker run --name postgres12_alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12_alpine createdb --username=root --owner=root facebook_db
	
dropdb:
	docker exec -it postgres12_alpine dropdb stats_db

migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/stats_db?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/stats_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

cert:
	cd cert; ./gen.sh; cd ..

build:
	@echo "[*] Building $(PROJECT_NAME)..."
	go build $(LDFLAGS) -o $(BIN_DIR)/$(BIN_FILE) $(CMD_DIR)/server/...
	@echo "[*] Finish..."


run_dev_grpc: export DATABASE_DEV_URL := postgresql://root:secret@localhost:5435/stats_db?sslmode=disable
run_dev_grpc: export SERVER_MODE = grpc
run_dev_grpc: export STATS_SERVICE_GRPC_SERVER_DIR = 0.0.0.0:50051
run_dev_grpc: export USER_SERVICE_GRPC_SERVER_DIR = 0.0.0.0:50052

run_dev_grpc:
	@echo "[*] Starting $(PROJECT_NAME)..."

	go run $(CMD_DIR)/server/...


run_dev_rest: export DATABASE_DEV_URL := postgresql://root:secret@localhost:5435/stats_db?sslmode=disable
run_dev_rest: export SERVER_MODE = rest
run_dev_rest: export GRPC_REST_ENDPOINT = 0.0.0.0:50051
run_dev_rest: export STATS_SERVICE_GRPC_SERVER_DIR = 0.0.0.0:3001
run_dev_rest: export USER_SERVICE_GRPC_SERVER_DIR = 0.0.0.0:50052
run_dev_rest: export ENABLE_TLS=YES
run_dev_rest:
	@echo "[*] Starting $(PROJECT_NAME)..."
	@echo "[*] Require that GRPC server is ON in port 50051..."

	go run $(CMD_DIR)/server/...



install:
	mkdir -p /etc/$(PROJECT_NAME)/
	cp $(BIN_DIR)/$(BIN_FILE) /usr/local/bin/
	echo $(shell ls)
	# cp $(CONFIG_DIR)/$(CONFIG_FILE).json /etc/

run:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test  build run_dev_rest install cert run