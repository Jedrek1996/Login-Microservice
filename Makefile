postgres:
	docker run --name postgres12 -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	docker exec -it postgres12 createdb --username=root --owner=root FoodPanda9

dropdb:
	docker exec -it postgres12 dropdb FoodPanda9

migrateup: 
	migrate -path $(PWD)/database/migration -database "postgresql://root:secret@localhost:5430/FoodPanda9?sslmode=disable" -verbose up

migratedown:
	migrate -path $(PWD)/database/migration -database "postgresql://root:secret@localhost:5430/FoodPanda9?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

#Build docker image
build:
	docker build -t login9 .

#Run docker image
run: 
	docker run -p 8080:8080 login9

##START SERVICE##
start:
	go run cmd/main.go 

.PHONY: createdb dropdb postgres migratedown migrateup test build run start
