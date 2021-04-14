test:
	go test -race ./...

build:
	@GOPROXY=https://goproxy.cn go build ./cmd/server
	@echo "Now can run server with ./server"

docker_build:
	docker run --rm -it -v "${PWD}:/app" -w /app golang:1.14 make build
