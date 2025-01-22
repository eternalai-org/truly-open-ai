build_osx:
	docker run --rm -v "$$(pwd)":/app -w /app golang:1.23 sh -c 'GOOS=darwin go build -o eai-chat main.go'

build_linux:
	docker run --rm -v "$$(pwd)":/app -w /app golang:1.23 sh -c 'GOOS=linux go build -o eai-chat main.go'
