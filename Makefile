SERVER_BINARY = ./bin/server
CLIENT_BINARY = ./bin/client

build-server:
	go build -o $(SERVER_BINARY) ./cmd/server/main.go

build-client:
	go build -o $(CLIENT_BINARY) ./cmd/client/main.go

run-server:
	$(SERVER_BINARY)

# PORT and POKEMON_LIST need to be defined by user during make call
run-client:
	$(CLIENT_BINARY) -port $(PORT) -pokemon $(POKEMON_LIST)

build-run-server: build-server run-server

build-run-client: build-client run-client

# make run-client PORT=8080 POKEMON_LIST=charizard,piplup,blastoise,pikachu,rayquaza,squirtle,bulbasaur,rhydon,charmander,mewtwo,mew,wartortle,charmeleon,salamence

# Prevents issues if files with these names exist
.PHONY: build-server build-client run-server run-client build-run-server build-run-client
