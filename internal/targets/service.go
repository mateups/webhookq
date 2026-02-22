package targets

import (
	"context"
	"database/sql"
	"errors"
	"net/url"

	"webhooq/internal/fault"
)

var (
	errTargetNotFound        = fault.NotFoundError("Target not found.")
	errInvalidTargetID       = fault.ValidationError("Target ID is required.")
	errInvalidTargetURL      = fault.ValidationError("Invalid target URL.")
	errInvalidRequestTimeout = fault.ValidationError("Timeout must be between 100 and 30000 ms.")
	errInvalidMaxAttempts    = fault.ValidationError("Max attempts must be between 1 and 20.")
	errFailedCreateTarget    = fault.InternalError("Failed to create target.")
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateTarget(ctx context.Context, input CreateTargetInput) (Target, error) {
	if err := validateCreateTargetInput(input); err != nil {
		return Target{}, err
	}

	target, err := s.repository.Create(ctx, input)
	if err != nil {
		return Target{}, errFailedCreateTarget
	}

	return target, nil
}

func (s *Service) ListTargets(ctx context.Context) ([]Target, error) {
	return s.repository.List(ctx)
}

func (s *Service) GetTarget(ctx context.Context, id string) (Target, error) {
	if id == "" {
		return Target{}, errInvalidTargetID
	}

	target, err := s.repository.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Target{}, errTargetNotFound
		}
		return Target{}, err
	}

	return target, nil
}

func validateCreateTargetInput(input CreateTargetInput) error {
	if input.Url == "" {
		return errInvalidTargetURL
	}

	parsedURL, err := url.Parse(input.Url)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errInvalidTargetURL
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errInvalidTargetURL
	}

	if input.RequestTimeoutMs < 100 || input.RequestTimeoutMs > 30000 {
		return errInvalidRequestTimeout
	}

	if input.MaxAttempts < 1 || input.MaxAttempts > 20 {
		return errInvalidMaxAttempts
	}

	return nil
}
