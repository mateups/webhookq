package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"webhooq/internal/fault"
)

type response struct {
	HttpStatus
	Payload any    `json:"payload,omitempty"`
	Error   string `json:"error,omitempty"`
}

func newJSONErrorResponse(responseWriter http.ResponseWriter, status HttpStatus, err string) {
	writeHeader(responseWriter, status.StatusCode)
	apiResp := response{
		HttpStatus: status,
		Error:      err,
	}
	_ = json.NewEncoder(responseWriter).Encode(apiResp)
}

func newJSONResponse(responseWriter http.ResponseWriter, payload any) {
	writeHeader(responseWriter, StatusOK.StatusCode)
	apiResp := response{
		HttpStatus: StatusOK,
		Payload:    payload,
	}
	_ = json.NewEncoder(responseWriter).Encode(apiResp)
}

func writeServiceError(responseWriter http.ResponseWriter, err error) {
	var appErr fault.Error
	if errors.As(err, &appErr) {
		var status HttpStatus
		switch appErr.Kind {
		case fault.KindValidation:
			status = StatusBadRequest
		case fault.KindNotFound:
			status = StatusNotFound
		default:
			status = StatusInternal
		}
		newJSONErrorResponse(responseWriter, status, appErr.Message)
		return
	}

	newJSONErrorResponse(responseWriter, StatusInternal, StatusInternal.Message)
}

func writeHeader(responseWriter http.ResponseWriter, status int) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
}
