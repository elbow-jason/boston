package learner

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	errorDeleteNameRequired = errors.New("DeleteAction 'name' cannot be blank")
)

// DeleteAction .
type DeleteAction struct {
	Name            string `json:"name"`
	SuccessCallback func()
	FailureCallback func(error)
}

// LearnerName .
func (a DeleteAction) LearnerName() string {
	return a.Name
}

// Validate a DeleteAction
func (a DeleteAction) Validate() error {
	if a.Name == "" {
		return errorDeleteNameRequired
	}
	return nil
}

// Signal .
func (a DeleteAction) Signal() int {
	return DeleteSignal
}

// Success a callback upon success
func (a DeleteAction) Success() {
	fmt.Println("DeleteAction Success Called")
	if a.SuccessCallback != nil {
		a.SuccessCallback()
	}
}

// Payload returns nil
func (a DeleteAction) Payload() interface{} {
	return nil
}

// Failure a callback upon failure
func (a DeleteAction) Failure(err error) {
	fmt.Println("DeleteAction Failure Called - ", err.Error())
	if a.FailureCallback != nil {
		a.FailureCallback(err)
	}
}

// DeleteActionFromJSON .
func DeleteActionFromJSON(r io.Reader) (DeleteAction, error) {
	action := DeleteAction{}
	err := json.NewDecoder(r).Decode(&action)
	if err != nil {
		return action, err
	}
	if err = action.Validate(); err != nil {
		return action, err
	}
	return action, nil
}
