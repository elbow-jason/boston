package main

import (
	"encoding/json"
	"log"
)

// APIError .
type APIError struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

// NewAPIError .
func NewAPIError(reason, message string) APIError {
	return APIError{
		Reason:  reason,
		Message: message,
	}
}

// APIErrors .
type APIErrors struct {
	Errors []APIError `json:"errors"`
}

func (a *APIErrors) addErrors(newErrors ...APIError) {
	a.Errors = append(a.Errors, newErrors...)
}

func (a *APIErrors) marshalJSON() []byte {
	bytes, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func newAPIErrors(errors ...APIError) *APIErrors {
	return &APIErrors{
		Errors: errors,
	}
}
