	# use --rm to remove the container once stop
	docker run --rm --name postgres12 -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine;
	
	# add sleep of 3 seconds, otherwise, too fast action might not have enough time for postgres container get ready
	sleep 3
	docker exec -it postgres12 createdb --username=root --owner=root loginMicroservice9;
	
	# same, need some time for the next step
	sleep 3
	migrate -path database/migration -database "postgresql://root:secret@localhost:5430/loginMicroservice9?sslmode=disable" -verbose up

	docker build -t login9 .;

	# use --rm to remove the container once stop
	docker run -p 8080:8080 login9


