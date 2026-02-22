package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const maxRequestBodyBytes = 1 << 20
const errorInvalidRequestBody = "Invalid request body"

func decodeJSON(responseWriter http.ResponseWriter, request *http.Request, target any) error {
	request.Body = http.MaxBytesReader(responseWriter, request.Body, maxRequestBodyBytes)

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(target); err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New(errorInvalidRequestBody)
		}
		return err
	}

	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New(errorInvalidRequestBody)
	}

	return nil
}
