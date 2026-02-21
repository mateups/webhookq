package targets

type Target struct {
	ID               string `json:"id"`
	URL              string `json:"url"`
	SigningSecret    string `json:"signing_secret,omitempty"`
	RequestTimeoutMS int    `json:"timeout_ms"`
	MaxAttempts      int    `json:"max_attempts"`
}

type CreateTargetInput struct {
	URL              string `json:"url"`
	SigningSecret    string `json:"signing_secret,omitempty"`
	RequestTimeoutMS int    `json:"timeout_ms"`
	MaxAttempts      int    `json:"max_attempts"`
}
