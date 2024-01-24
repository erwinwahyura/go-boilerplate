package model

type (
	// HealthCheckResponse response
	HealthCheckResponse struct {
		Success []string `json:"success,omitempty"`
		Errors  []string `json:"errors,omitempty"`
	}
)
