package api

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"webhooq/internal/fault"
)

type responseEnvelope struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func TestWriteServiceError_ValidationFault(t *testing.T) {
	recorder := httptest.NewRecorder()

	writeServiceError(recorder, fault.ValidationError("bad input"))

	var envelope responseEnvelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if recorder.Code != 400 {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}
	if envelope.Error != "bad input" {
		t.Fatalf("expected validation message, got %q", envelope.Error)
	}
}

func TestWriteServiceError_NotFoundFault(t *testing.T) {
	recorder := httptest.NewRecorder()

	writeServiceError(recorder, fault.NotFoundError("missing"))

	if recorder.Code != 404 {
		t.Fatalf("expected 404, got %d", recorder.Code)
	}
}

func TestWriteServiceError_UnknownError(t *testing.T) {
	recorder := httptest.NewRecorder()

	writeServiceError(recorder, errors.New("raw db error"))

	var envelope responseEnvelope
	if err := json.Unmarshal(recorder.Body.Bytes(), &envelope); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if recorder.Code != 500 {
		t.Fatalf("expected 500, got %d", recorder.Code)
	}
	if envelope.Error != StatusInternal.Message {
		t.Fatalf("expected internal message %q, got %q", StatusInternal.Message, envelope.Error)
	}
}
