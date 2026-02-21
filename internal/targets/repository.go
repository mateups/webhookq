package targets

import "context"

type Repository interface {
	Create(ctx context.Context, input CreateTargetInput) (Target, error)
	Get(ctx context.Context, targetID string) (Target, error)
	List(ctx context.Context) ([]Target, error)
}
