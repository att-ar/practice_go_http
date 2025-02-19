package common

// pokemonJSON is a struct for defining parsing of the JSON response
type PokemonJSON struct {
	Name  string     `json:"name"` // Maps the JSON field "name" to the struct field Name.
	Types []struct { // A slice of anonymous structs.
		Type struct { // An anonymous struct inside the slice.
			Name string `json:"name"` // Maps the JSON field "name" inside "type" to the struct field Name.
		} `json:"type"` // Maps the JSON field "type" to the struct field Type.
	} `json:"types"` // Maps the JSON field "types" to the struct field Types.
}

// Pokemon is a struct to hold the data we care about
type Pokemon struct {
	Name string
	Type []string
}
