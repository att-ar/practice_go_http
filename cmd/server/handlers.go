package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"practice_http/cmd/common"

	"github.com/gorilla/mux"
)

func newHandler() http.Handler {
	router := mux.NewRouter()

	// pokemon
	router.HandleFunc("/pokemon", handlePostPokemon).Methods("POST")

	return router
}

/* Print the data of Pokemons that are posted to this endpoint */
func handlePostPokemon(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	// Attempt to decode as an array of Pokemons
	if _, err := decoder.Token(); err != nil { // Check for the start of the array
		http.Error(rw, "Expected '['", http.StatusBadRequest)
		return
	}

	// Loop through the pokemons now that they're at the top level of the decoder
	var pokemon common.Pokemon
	for decoder.More() {
		err := decoder.Decode(&pokemon)
		if err != nil {
			log.Printf("Decoding Error: %v\n", err)            // log the error to the server console
			http.Error(rw, err.Error(), http.StatusBadRequest) // send 400 back to client
			return
		}

		log.Printf("Pokemon received: %+v\n", pokemon) // "store" the pokemon lol
	}

	if _, err := decoder.Token(); err != nil { // Check for the end of the array
		http.Error(rw, "Expected ']'", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(rw, "Pokemons processed") // Success message back to client
}
