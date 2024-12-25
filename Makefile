.PHONY: help
help:
	@echo "Commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:
	@go build -o bin/papestash cmd/main.go

.PHONY: run
run: build
	@./bin/papestash

.PHONY: migrate
migrate:
	sqlite3 db.database < migrations/up.sql
