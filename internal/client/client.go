/* Handle the PokeAPI */
package client

import (
	"context"
	"encoding/json"
	"io"
	c "practice_http/internal/common"
	"practice_http/pkg/httpclient"
	"time"
)

const BASE_URL = "https://pokeapi.co/api/v2/pokemon/"

func GetPokemon(name string) (*c.Pokemon, error) {
	url := BASE_URL + name

	// 8 second timeout is probably enough
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	// httpclient.GET[*c.Pokemon] is inferred from decodePokemon
	pokemon, err := httpclient.GET(ctx, url, decodePokemon)
	if err != nil {
		var emptyPokemon *c.Pokemon
		return emptyPokemon, err
	}

	return pokemon, nil
}

func decodePokemon(body io.Reader) (*c.Pokemon, error) {
	// pokemonJSON target to decode into:
	pokemonJSON := &c.PokemonJSON{}
	// decode into target (not using unmarshal because we get io.ReadCloser from the response)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(pokemonJSON)
	if err != nil {
		var emptyPokemon *c.Pokemon
		return emptyPokemon, err
	}

	// Convert pokemonJSON to Pokemon
	pokemon := &c.Pokemon{
		Name: pokemonJSON.Name,
		Type: make([]string, len(pokemonJSON.Types)),
	}
	// Simplify the Types field to just a slice of strings for Pokemon types
	for i, t := range pokemonJSON.Types {
		pokemon.Type[i] = t.Type.Name
	}

	return pokemon, nil
}
