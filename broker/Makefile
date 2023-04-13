build:
	- go build -o broker src/main.go

run:
	- go run src/main.go

docker-build:
	- docker build . -t guilhermeroque/broker

docker-run-bash:
	- docker run -it guilhermeroque/broker sh

docker-run-default:
	- docker run -it -p 8081:8081 guilhermeroque/broker

docker-run-detached:
	- docker run -d guilhermeroque/broker

