.PHONY: build run logs dockerstop
.SILENT: build run logs dockerstop

run:
	docker compose build
	docker compose up

install:
	go mod tidy