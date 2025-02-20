package httpclient

import (
	"context"
	"io"
	"log"
	"net/http"
)

/*
Generic Interface to dependency inject json.Decode logic without worrying about closing the response.
*/
type Decoder[T any] func(io.Reader) (T, error)

/*
Returns the return value of the Decoder if the GET request works properly
So the error will be what you define in the Decoder.
*/
func GET[T any](ctx context.Context, url string, d Decoder[T]) (T, error) {
	// create GET request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		var zero T // zero value of T
		return zero, err
	}
	req.Header.Add("Accept", "application/json") // only dealing with json

	// send GET request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		var zero T
		return zero, err
	}
	defer res.Body.Close()

	// return the result of the Decoder
	return d(res.Body)
}

func POST(ctx context.Context, url string, body io.Reader) error {
	// create POST request
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json") // only dealing with json

	// send POST request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		// Read error body if there is one
		errorBody, _ := io.ReadAll(res.Body)
		log.Fatalf("Unexpected status code: %d. Error: %s\n", res.StatusCode, string(errorBody))
	}
	// Read the response body as a string
	resBody, err := io.ReadAll(res.Body) // Read the response body
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	log.Printf("Server response: %s", string(resBody))
	return nil
}
