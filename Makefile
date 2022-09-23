postgres:
	sudo docker run --name postgres12 -p 8081:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root pathshaala

dropdb:
	sudo docker exec -it postgres12 dropdb pathshaala

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:8081/pathshaala?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:8081/pathshaala?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc