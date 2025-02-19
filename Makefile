SERVER_BINARY = ./bin/server
CLIENT_BINARY = ./bin/client

build-server:
	go build -o $(SERVER_BINARY) ./cmd/server/main.go

build-client:
	go build -o $(CLIENT_BINARY) ./cmd/client/main.go

run-server:
	$(SERVER_BINARY)

run-client:
	$(CLIENT_BINARY) -server-address localhost:8080

build-run-server: build-server run-server

build-run-client: build-client run-client

# Prevents issues if files with these names exist
.PHONY: build-server build-client run-server run-client build-run-server build-run-client
