/* Handle the server */
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	c "practice_http/internal/common"
	"practice_http/pkg/httpclient"
	"time"
)

func PostPokemon(pokemon *c.Pokemon, url string) error {
	// marshal using the json field names declared in common/types.go
	marshalled, err := json.Marshal(pokemon)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	// make the io.Reader
	reader := bytes.NewReader(marshalled)

	// send POST
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	err = httpclient.POST(ctx, url, reader)
	if err != nil {
		return err
	}
	return nil
}
