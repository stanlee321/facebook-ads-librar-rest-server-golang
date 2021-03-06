PROJECT_NAME = microservice-ads
BIN_DIR = ./bin
BIN_FILE = $(PROJECT_NAME)
CMD_DIR = ./cmd

# Get version constant
VERSION := 1.0.0
BUILD := $(shell git rev-parse HEAD)

# Use linker flags to provide version/build settings to the binary
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION) -X=main.build=$(BUILD)"

# Some ENV VARS

gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pkg/api/v1 --grpc-gateway_out=:pkg/api/v1 --openapiv2_out=:openapiv2

postgres:
	docker run --name postgres12_alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres12_alpine createdb --username=root --owner=root facebook_ads
	
dropdb:
	docker exec -it postgres12_alpine dropdb facebook_ads

migrateinit:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/facebook_ads?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/facebook_ads?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

cert:
	cd cert; ./gen.sh; cd ..

build:
	@echo "[*] Building $(PROJECT_NAME)..."
	go build $(LDFLAGS) -o $(BIN_DIR)/$(BIN_FILE)
	@echo "[*] Finish..."

build_docker_image:
	docker build --no-cache -t stanlee321/ads/facebook_ads:latest -f ./Dockerfile .

tag:
	docker tag stanlee321/ads/facebook_ads:latest stanlee321/facebook_ads_api:latest

publish:
	docker push stanlee321/facebook_ads_api:latest

install:
	mkdir -p /etc/$(PROJECT_NAME)/
	cp $(BIN_DIR)/$(BIN_FILE) /usr/local/bin/
	echo $(shell ls)
	# cp $(CONFIG_DIR)/$(CONFIG_FILE).json /etc/

run:
	go run main.go

up:
	docker volume create --name=db-psql-db
	docker-compose -f docker-compose.yml up --build  -d
	sleep 5
	docker-compose exec web python manage.py db init
	docker-compose exec web python manage.py db migrate
	docker-compose exec web python manage.py db upgrade

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test  build run_dev_rest install cert run