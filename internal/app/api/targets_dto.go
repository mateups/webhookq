package api

// Request DTOs
type createTargetRequest struct {
	Url              string `json:"url"`
	SigningSecret    string `json:"signing_secret,omitempty"`
	RequestTimeoutMs int    `json:"request_timeout_ms"`
	MaxAttempts      int    `json:"max_attempts"`
}

// Response DTOs
type targetResponse struct {
	Id               string `json:"id"`
	Url              string `json:"url"`
	RequestTimeoutMs int    `json:"request_timeout_ms"`
	MaxAttempts      int    `json:"max_attempts"`
}

type listTargetsResponse struct {
	Items []targetResponse `json:"items"`
}
