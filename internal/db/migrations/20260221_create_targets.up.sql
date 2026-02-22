CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE SCHEMA IF NOT EXISTS webhookq;

CREATE TABLE webhookq.targets(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    url text NOT NULL,
    signing_secret text NULL,
    request_timeout_ms int NOT NULL,
    CHECK (request_timeout_ms >= 100),
    CHECK (request_timeout_ms <= 30000),
    max_attempts int NOT NULL,
    CHECK (max_attempts >= 1),
    CHECK (max_attempts <= 20),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
)