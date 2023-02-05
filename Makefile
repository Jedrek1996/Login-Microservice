postgres:
	docker run --name postgres12 -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it postgres12 createdb --username=root --owner=root FoodPanda9

dropdb:
	docker exec -it postgres12 dropdb FoodPanda9

migrateup: 
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/FoodPanda9?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5432/FoodPanda9?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migratedown migrateup test
