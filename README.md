# webhookq
A self-hostable, serverless-style webhook delivery service: an API to enqueue delayed jobs and schedules, plus a worker that reliably delivers JSON payloads to registered HTTP endpoints with retries and observability. The system is language-agnostic: it does not execute user code—“execution” is an outbound HTTP POST to a target URL.
