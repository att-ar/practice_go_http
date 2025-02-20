package main

import (
	"log"
	"practice_http/internal/client"
)

func main() {
	pokemon, err := client.GetPokemon("charizard")
	if err != nil {
		log.Println("Failed to retrieve pokemon: ", err.Error())
	}
	log.Printf("Pokemon retrieved: %+v\n", pokemon)
}
