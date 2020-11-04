
read:
	go run ./cmd/*.go read

server:
	go run ./cmd/*.go server

build:
	go build -o manager ./cmd/**.go

docker-build:
	docker build -t logapi .

docker-run:
	docker run -it -v ${PWD}/examples:/examples -p 8080:8080 logapi server

curl:
	curl 127.0.0.1:8080/files

stress:
	siege -v -t 1 -c 100 http://127.0.0.1:8080/files