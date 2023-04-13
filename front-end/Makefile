build:
	- go build -o front-end-broker ./cmd/web/main.go

run:
	- go run ./cmd/web/main.go

docker-build:
	- docker build . -t guilhermeroque/front-end-broker

docker-run-bash:
	- docker run -it guilhermeroque/front-end-broker sh

docker-run-default:
	- docker run -it -p 8080:8080 guilhermeroque/front-end-broker 
docker-run-detached:
	- docker run -d guilhermeroque/front-end-broker

