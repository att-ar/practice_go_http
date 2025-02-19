server:
	cd cmd/server && go build

client:
	cd cmd/client && go build

run-server:
	cd cmd/server && ./server

run-client:
	cd cmd/client && ./client -server-address localhost:8080

build-run-server:
	make server && ./server

build-run-client:
	make client && ./client -server-address localhost:8080

.PHONY: server client run-server run-client # Prevents issues if files with these names exist
