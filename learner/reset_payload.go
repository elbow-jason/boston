package learner

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	errorResetNameRequired = errors.New("Reset 'name' cannot be blank")
)

// ResetPayload .
type ResetPayload struct {
	Name string `json:"name"`
}

// ValidateResetPayload .
func ValidateResetPayload(payload ResetPayload) error {
	if payload.Name == "" {
		return errorResetNameRequired
	}
	return nil
}

// ResetPayloadFromJSON .
func ResetPayloadFromJSON(r io.Reader) (ResetPayload, error) {
	payload := ResetPayload{}
	err := json.NewDecoder(r).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, ValidateResetPayload(payload)
}
