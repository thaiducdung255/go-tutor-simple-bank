migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose down

psql:
	docker pull postgres:12-alpine
	docker run --name pg12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it pg12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it pg12 dropdb simple_bank

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb
