name=postgres15
dbport=5433
dbhost=127.0.0.1

rundb:
	docker run --name ${name} -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d postgres:15.2-alpine

shelldb:
	docker exec -it ${name} psql --username=postgres --db=users

createdb:
	docker exec -it ${name} createdb --username=postgres --owner=postgres users

dropdb:
	docker exec -it ${name} dropdb --username=postgres users

migratecreate:
	migrate create -ext sql -dir ./infra/migrations -seq $(name)

migrateup:
	migrate -database postgres://postgres:password@${dbhost}:${dbport}/users?sslmode=disable -path ./infra/migrations -verbose up

migratedown:
	migrate -database postgres://postgres:password@${dbhost}:${dbport}/users?sslmode=disable -path ./infra/migrations -verbose down


