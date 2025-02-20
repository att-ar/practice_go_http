package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"practice_http/internal/common"

	"github.com/gorilla/mux"
)

func newHandler() http.Handler {
	router := mux.NewRouter()

	// /
	router.HandleFunc("/", handleHandshake).Methods("GET", "POST")

	// /pokemon
	router.HandleFunc("/pokemon", handlePostPokemon).Methods("POST")

	// /pokemon/list
	router.HandleFunc("/pokemon/list", handlePostPokemonList).Methods("POST")

	return router
}

/* Establish handshake */
func handleHandshake(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	fmt.Fprintf(rw, "Handshake Successful\n") // Success message back to client
}

/* Print the data of Pokemons posted to this endpoint */
func handlePostPokemon(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)

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
	fmt.Fprintf(rw, "Pokemon processed\n") // Success message back to client
}

/* Print the data of Pokemons that are in a list posted to this endpoint */
func handlePostPokemonList(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)

	// Attempt to decode as an array of Pokemons
	if _, err := decoder.Token(); err != nil { // Check for the start of the array
		http.Error(rw, "Expected '['", http.StatusBadRequest)
		return
	}

	// Loop through the pokemons now that they're at the top level of the decoder
	var pokemon common.Pokemon
	count := 0
	for decoder.More() {
		err := decoder.Decode(&pokemon)
		if err != nil {
			log.Printf("Decoding Error: %v\n", err)            // log the error to the server console
			http.Error(rw, err.Error(), http.StatusBadRequest) // send 400 back to client
			return
		}

		log.Printf("Pokemon received: %+v\n", pokemon) // "store" the pokemon lol
		count += 1
	}

	if _, err := decoder.Token(); err != nil { // Check for the end of the array
		http.Error(rw, "Expected ']'", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "%d Pokemons processed\n", count) // Success message back to client
}
