package httpclient

import (
	"context"
	"io"
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
