package targets

import (
	"context"
	"database/sql"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

// Create implements [Repository].
func (r *postgresRepository) Create(ctx context.Context, input CreateTargetInput) (Target, error) {
	target := Target{}

	query := `
		INSERT INTO webhookq.targets (url, signing_secret, request_timeout_ms, max_attempts)
		VALUES ($1, $2, $3, $4)
		RETURNING id, url, signing_secret, request_timeout_ms, max_attempts
	`
	err := r.db.QueryRowContext(ctx, query, input.Url, input.SigningSecret, input.RequestTimeoutMs, input.MaxAttempts).Scan(
		&target.Id,
		&target.Url,
		&target.SigningSecret,
		&target.RequestTimeoutMs,
		&target.MaxAttempts,
	)

	if err != nil {
		return Target{}, err
	}

	return target, nil
}

// Get implements [Repository].
func (r *postgresRepository) Get(ctx context.Context, targetId string) (Target, error) {
	target := Target{}

	query := `
		SELECT id, url, signing_secret, request_timeout_ms, max_attempts
		FROM webhookq.targets
		WHERE id = $1
	`
	err := r.db.QueryRowContext(ctx, query, targetId).Scan(&target.Id, &target.Url, &target.SigningSecret, &target.RequestTimeoutMs, &target.MaxAttempts)

	if err == sql.ErrNoRows {
		return Target{}, err
	}

	if err != nil {
		return Target{}, err
	}

	return target, nil
}

// List implements [Repository].
func (r *postgresRepository) List(ctx context.Context) ([]Target, error) {
	query := `
		SELECT id, url, signing_secret, request_timeout_ms, max_attempts
		FROM webhookq.targets
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var targets []Target
	for rows.Next() {
		var target Target
		err := rows.Scan(&target.Id, &target.Url, &target.SigningSecret, &target.RequestTimeoutMs, &target.MaxAttempts)
		if err != nil {
			return nil, err
		}
		targets = append(targets, target)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return targets, nil
}
