package coc

import "fmt"

type APIError struct {
	StatusCode int
	Reason     string `json:"reason,omitempty"`
	Message    string `json:"message,omitempty"`
	Err        error
}

func (a *APIError) Error() string {
	return fmt.Sprintf("[%d] Reason: %s Message: %s", a.StatusCode, a.Reason, a.Message)
}

// Expand upon this later...
var (
	BadRequest           = "badRequest"
	InvalidAuthorization = "accessDenied"
	InvalidIP            = "accessDenied.invalidIp"
	NotFound             = "notFound"
)
