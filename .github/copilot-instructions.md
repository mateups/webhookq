# Agent Instructions for `webhookq`

This repo favors **small, explicit, layered** changes. Optimize for correctness and maintainability over novelty.

## Non-Negotiables

- Ship the **smallest correct fix** for the asked scope.
- Do not add extra layers/interfaces/abstractions unless explicitly requested or removing active duplication.
- Keep unrelated behavior and files unchanged.

## Architecture Rules

- HTTP/transport concerns stay in API layer.
- Business rules + input validation stay in domain services.
- SQL/persistence stays in repositories.
- Domain code must not depend on HTTP status/types.

## Error Strategy (single pattern)

- Define reusable domain faults once in a shared package.
- Map domain faults to HTTP responses in one consistent path.
- Unknown/untyped failures => generic `500 Internal server error`.
- Never return raw DB/repository internals to clients.

## API/JSON Rules

- Use standard library `net/http` by default.
- Keep payloads simple and consistent.
- Use `omitempty` only for truly optional fields.
- Build response payloads in transport/handler code (not DTO declarations).

## Validation/Data Rules

- Validate in domain service (JSON tags alone are not validation).
- Keep migrations, schema names, and repository SQL aligned.
- DB constraints are safety net; service should return user-friendly validation errors first.

## Quality Bar

- Keep naming/messages consistent (prefer `Id`/`Ms` style, stable wording).
- Prefer readability and directness over clever code.
- After changes, run `go build ./...`; if tests exist for touched areas, run targeted tests.
