package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"practice_http/internal/client"
	c "practice_http/internal/common"
	"strings"
	"sync"
)

func main() {
	// command-line flags for the port (my own server) and the pokemons to fetch and post to the server
	port := flag.String("port", "8080", "Port to connect to")
	pokemonList := flag.String("pokemon", "", "Comma-separated list of Pokemon names")
	flag.Parse()

	// URLs:
	SERVER_URL := fmt.Sprintf("http://localhost:%s", *port)
	SERVER_POKEMON_URL := fmt.Sprintf("%s/pokemon", SERVER_URL)

	// handshake with server
	res, err := http.Get(SERVER_URL)
	if err != nil {
		log.Printf("Failed to handshake with server at port %s\n", *port)
		return
	}
	defer res.Body.Close()
	log.Printf("Completed handshake with server at port %s\n", *port)

	// Split the list
	pokemonNames := strings.Split(*pokemonList, ",")

	// will use this channel to grab the pokemon from the goroutines and post them to my server
	ch := make(chan *c.Pokemon)

	// get one by one
	wgFetch := new(sync.WaitGroup) // will handle waiting for fetching
	getPokemon := func(pokemonName string, wg *sync.WaitGroup, ch chan<- *c.Pokemon) {
		defer wg.Done()
		pokemon, err := client.GetPokemon(pokemonName)
		if err != nil {
			log.Println("Failed to retrieve pokemon: ", err.Error())
		}
		ch <- pokemon
		// log.Printf("Pokemon retrieved: %+v\n", pokemon)
	}
	for _, pokemonName := range pokemonNames {
		wgFetch.Add(1)
		go getPokemon(pokemonName, wgFetch, ch)
	}

	// spawn goroutine that handles closing the channel so that the wait isn't blocking
	go func() {
		wgFetch.Wait()
		close(ch) // close the channel (blocks writing not reading)
	}()

	// post one by one
	wgPost := new(sync.WaitGroup) // will handle waiting for posting
	postPokemon := func(pokemon *c.Pokemon, wg *sync.WaitGroup) {
		defer wg.Done()
		err := client.PostPokemon(pokemon, SERVER_POKEMON_URL)
		if err != nil {
			log.Printf("Failed to post pokemon: %+v\n", pokemon)
		}
		// log.Println("Pokemon posted.")
	}
	for pokemon := range ch {
		wgPost.Add(1)
		go postPokemon(pokemon, wgPost)
	}

	// wait for everything to finish posting
	wgPost.Wait()
}
