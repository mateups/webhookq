package api

import (
	"net/http"

	"webhooq/internal/targets"
)

type TargetsHandler struct {
	service *targets.Service
}

func NewTargetsHandler(service *targets.Service) *TargetsHandler {
	return &TargetsHandler{service: service}
}

func (handler *TargetsHandler) CreateTarget(responseWriter http.ResponseWriter, request *http.Request) {
	var req createTargetRequest
	if err := decodeJSON(responseWriter, request, &req); err != nil {
		newJSONErrorResponse(responseWriter, StatusBadRequest, "Invalid request body")
		return
	}

	createInput := targets.CreateTargetInput{
		Url:              req.Url,
		SigningSecret:    req.SigningSecret,
		RequestTimeoutMs: req.RequestTimeoutMs,
		MaxAttempts:      req.MaxAttempts,
	}

	createdTarget, err := handler.service.CreateTarget(request.Context(), createInput)
	if err != nil {
		writeServiceError(responseWriter, err)
		return
	}

	response := targetResponse{
		Id:               createdTarget.Id,
		Url:              createdTarget.Url,
		RequestTimeoutMs: createdTarget.RequestTimeoutMs,
		MaxAttempts:      createdTarget.MaxAttempts,
	}
	newJSONResponse(responseWriter, response)
}

func (handler *TargetsHandler) ListTargets(responseWriter http.ResponseWriter, request *http.Request) {
	result, err := handler.service.ListTargets(request.Context())
	if err != nil {
		writeServiceError(responseWriter, err)
		return
	}

	items := make([]targetResponse, 0, len(result))
	for _, target := range result {
		items = append(items, targetResponse{
			Id:               target.Id,
			Url:              target.Url,
			RequestTimeoutMs: target.RequestTimeoutMs,
			MaxAttempts:      target.MaxAttempts,
		})
	}
	response := listTargetsResponse{Items: items}
	newJSONResponse(responseWriter, response)
}

func (handler *TargetsHandler) GetTarget(responseWriter http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")

	if id == "" {
		newJSONErrorResponse(responseWriter, StatusBadRequest, "Target ID is required.")
		return
	}

	result, err := handler.service.GetTarget(request.Context(), id)
	if err != nil {
		writeServiceError(responseWriter, err)
		return
	}

	response := targetResponse{
		Id:               result.Id,
		Url:              result.Url,
		RequestTimeoutMs: result.RequestTimeoutMs,
		MaxAttempts:      result.MaxAttempts,
	}
	newJSONResponse(responseWriter, response)
}
