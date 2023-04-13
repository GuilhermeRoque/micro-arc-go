up:
	- docker-compose up
build-up:
	- docker-compose up --build

up-detached:
	- docker-compose up -d
down:
	- docker-compose down -v
	- docker-compose rm

createdb-and-run-migrations:
	- docker exec -it micro-arc-go_postgres_1 createdb --username=postgres --owner=postgres users
	- migrate -database postgres://postgres:password@127.0.0.1:5433/users?sslmode=disable -path ../auth-service/infra/migrations -verbose up

up-migrations:
	- migrate -database postgres://postgres:password@127.0.0.1:5433/users?sslmode=disable -path ../auth-service/infra/migrations -verbose up

down-migrations:
	- migrate -database postgres://postgres:password@127.0.0.1:5433/users?sslmode=disable -path ../auth-service/infra/migrations -verbose down