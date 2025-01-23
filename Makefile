build_decentralize_server_osx:
	cd decentralized-inference && docker run --rm -v "$$(pwd)":/app -w /app golang:1.23 sh -c 'GOOS=darwin go build -o eai-chat main.go' \
	&& mv eai-chat ../eai-chat

build_decentralize_server_linux:
	cd decentralized-inference && docker run --rm -v "$$(pwd)":/app -w /app golang:1.23 sh -c 'GOOS=linux go build -o eai-chat main.go' \
	&& mv eai-chat ../eai-chat
