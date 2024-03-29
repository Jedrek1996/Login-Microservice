#include .env if exists
-include .env

tools:
	brew install golang-migrate

#Create postgres image and runs it
postgres:
	docker run --name postgres12 -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

#Creates Login Database
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root loginMicroservice9

#Deletes Login Datase
dropdb:
	docker exec -it postgres12 dropdb loginMicroservice9

#Creats the tables based on the schemas
migrateup:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5430/loginMicroservice9?sslmode=disable" -verbose up

#Deletes the tables based on the schemas
migratedown:
	migrate -path database/migration -database "postgresql://root:secret@localhost:5430/loginMicroservice9?sslmode=disable" -verbose down

#Generate test files
test:
	go test -v -cover ./...

#Generate golang queries methods with SQLC (External Library)
sqlc:
	sqlc generate

#Build Docker image for Login Service
build:
	docker build -t login9 .

#💥 FOR THE ABOVE ONLY GENERATE WHEN NECESSARY #💥
#In order to run this service open 2 CLI, one for each command below and make run, make start in each CLI.

#Run Login Service container (Need to run this in order to use in concurrently with other services)
run:
	docker run -p 5430:8080 login9

#docker start container 
#START SERVICE#🎋 (Should see a pop, click allow)
start:
	go run cmd/main.go

down:
	-docker stop postgres12
	-docker rm -f postgres12
	-docker stop login-service
	-docker rm -f login-service

start-container:
	-docker stop postgres12
	-docker rm -f postgres12
	-docker stop login-service
	-docker rm -f login-service
	-docker network rm login-service-net
	bash start.sh
	
start-host:
	-docker stop postgres12
	-docker rm -f postgres12
	-docker stop login-service
	-docker rm -f login-service
	bash start_host.sh

start-auto:
	@if ("$(Run_ON_HOST)" == "true"); then\
        echo "Hello world"; \
		make start-host; \
	else \
		make start-container; \
    fi



.PHONY: createdb dropdb postgres migratedown migrateup test build run start