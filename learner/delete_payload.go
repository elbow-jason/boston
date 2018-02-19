package learner

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	errorDeleteNameRequired = errors.New("Delete 'name' cannot be blank")
)

// DeletePayload is the payload to
// execute a delete operation on a learner.
type DeletePayload struct {
	Name string `json:"name"`
}

// ValidateDeletePayload .
func ValidateDeletePayload(payload DeletePayload) error {
	if payload.Name == "" {
		return errorDeleteNameRequired
	}
	return nil
}

// DeletePayloadFromJSON .
func DeletePayloadFromJSON(r io.Reader) (DeletePayload, error) {
	payload := DeletePayload{}
	err := json.NewDecoder(r).Decode(&payload)
	if err != nil {
		return payload, err
	}
	if err = ValidateDeletePayload(payload); err != nil {
		return payload, err
	}
	return payload, nil
}
