include .env
export $(shell sed 's/=.*//' .env)

migrate-up:
	GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	GOOSE_TABLE=$(GOOSE_TABLE) \
	go run github.com/pressly/goose/v3/cmd/goose@latest up

seed:
	go run tools/postgres/seeds/main.go

migrate-reset:
	GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=tools/postgres/migrations \
	GOOSE_TABLE=$(GOOSE_TABLE) \
	go run github.com/pressly/goose/v3/cmd/goose@latest reset