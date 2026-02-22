package targets

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"webhooq/internal/fault"
)

type stubRepository struct {
	createFn func(ctx context.Context, input CreateTargetInput) (Target, error)
	getFn    func(ctx context.Context, targetID string) (Target, error)
	listFn   func(ctx context.Context) ([]Target, error)
}

func (repository stubRepository) Create(ctx context.Context, input CreateTargetInput) (Target, error) {
	if repository.createFn != nil {
		return repository.createFn(ctx, input)
	}
	return Target{}, nil
}

func (repository stubRepository) Get(ctx context.Context, targetID string) (Target, error) {
	if repository.getFn != nil {
		return repository.getFn(ctx, targetID)
	}
	return Target{}, nil
}

func (repository stubRepository) List(ctx context.Context) ([]Target, error) {
	if repository.listFn != nil {
		return repository.listFn(ctx)
	}
	return nil, nil
}

func TestValidateCreateTargetInput(t *testing.T) {
	tests := []struct {
		name    string
		input   CreateTargetInput
		wantErr error
	}{
		{
			name: "valid input",
			input: CreateTargetInput{
				Url:              "https://example.com/webhook",
				RequestTimeoutMs: 1000,
				MaxAttempts:      3,
			},
		},
		{
			name: "invalid URL",
			input: CreateTargetInput{
				Url:              "not-a-url",
				RequestTimeoutMs: 1000,
				MaxAttempts:      3,
			},
			wantErr: errInvalidTargetURL,
		},
		{
			name: "invalid timeout",
			input: CreateTargetInput{
				Url:              "https://example.com/webhook",
				RequestTimeoutMs: 50,
				MaxAttempts:      3,
			},
			wantErr: errInvalidRequestTimeout,
		},
		{
			name: "invalid max attempts",
			input: CreateTargetInput{
				Url:              "https://example.com/webhook",
				RequestTimeoutMs: 1000,
				MaxAttempts:      0,
			},
			wantErr: errInvalidMaxAttempts,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateCreateTargetInput(test.input)
			if test.wantErr == nil && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if test.wantErr != nil && !errors.Is(err, test.wantErr) {
				t.Fatalf("expected error %v, got %v", test.wantErr, err)
			}
		})
	}
}

func TestCreateTargetRepositoryFailureReturnsInternalFault(t *testing.T) {
	service := NewService(stubRepository{
		createFn: func(ctx context.Context, input CreateTargetInput) (Target, error) {
			return Target{}, errors.New("db down")
		},
	})

	_, err := service.CreateTarget(context.Background(), CreateTargetInput{
		Url:              "https://example.com/webhook",
		RequestTimeoutMs: 1000,
		MaxAttempts:      3,
	})

	if !errors.Is(err, errFailedCreateTarget) {
		t.Fatalf("expected %v, got %v", errFailedCreateTarget, err)
	}

	var faultErr fault.Error
	if !errors.As(err, &faultErr) || faultErr.Kind != fault.KindInternal {
		t.Fatalf("expected internal fault kind, got %v", err)
	}
}

func TestGetTargetNotFoundMapsToDomainFault(t *testing.T) {
	service := NewService(stubRepository{
		getFn: func(ctx context.Context, targetID string) (Target, error) {
			return Target{}, sql.ErrNoRows
		},
	})

	_, err := service.GetTarget(context.Background(), "missing-id")
	if !errors.Is(err, errTargetNotFound) {
		t.Fatalf("expected %v, got %v", errTargetNotFound, err)
	}
}
