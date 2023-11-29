package servemux

import (
	"encoding/json"
	"io"
)

func Unmarshal(rc io.ReadCloser, v interface{}) error {
	// read request body
	xb, err := io.ReadAll(rc) // ServeHTTP closes the body
	if err != nil {
		return err
	}

	err = json.Unmarshal(xb, v)
	return err
}
