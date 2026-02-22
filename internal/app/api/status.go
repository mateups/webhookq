package api

import "net/http"

type HttpStatus struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
}

var (
	StatusOK         = HttpStatus{StatusCode: http.StatusOK, Success: true, Message: "OK"}
	StatusBadRequest = HttpStatus{StatusCode: http.StatusBadRequest, Success: false, Message: "Invalid request."}
	StatusNotFound   = HttpStatus{StatusCode: http.StatusNotFound, Success: false, Message: "Resource not found."}
	StatusInternal   = HttpStatus{StatusCode: http.StatusInternalServerError, Success: false, Message: "Internal server error."}
)
