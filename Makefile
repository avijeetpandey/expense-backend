createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres expense_tracker

dropdb:
	docker exec -it postgres dropdb  --username=postgres  expense_tracker

migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/expense_tracker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/expense_tracker?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc server