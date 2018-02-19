package learner

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGen = rand.New(randSource)

var (
	errorTrainNameRequired   = errors.New("Train 'name' cannot be blank")
	errorTrainEmptyEntries   = errors.New("Train 'entries' cannot be empty")
	errorTrainZeroInputsSize = errors.New("Train 'InputsSize' cannot be 0")
	errorTrainZeroLabelsSize = errors.New("Train 'LabelsSize' cannot be 0")
	errorTrainSplitTestRange = errors.New("Train 'split_test' must be between 0.0 and 1.0")
)

// TrainPayload .
type TrainPayload struct {
	Name        string      `json:"name"`
	DataEntries []DataEntry `json:"entries"`
	TestSplit   float64     `json:"test_split"`
	InputsSize  int
	LabelsSize  int
}

func applyTrainTestSplit(payload TrainPayload) ([]DataEntry, []DataEntry) {
	split := payload.TestSplit
	if split <= 0.0 {
		return payload.DataEntries, []DataEntry{}
	}
	trains := []DataEntry{}
	tests := []DataEntry{}
	for _, entry := range payload.DataEntries {
		if randGen.Float64() < split {
			tests = append(tests, entry)
		} else {
			trains = append(trains, entry)
		}
	}
	return trains, tests
}

// ValidateTrainPayload a TrainPayload's values
func ValidateTrainPayload(payload TrainPayload) error {
	inputsSize := payload.InputsSize
	labelsSize := payload.LabelsSize
	if inputsSize == 0 {
		return errorTrainZeroInputsSize
	}
	if labelsSize == 0 {
		return errorTrainZeroLabelsSize
	}
	if len(payload.DataEntries) == 0 {
		return errorTrainEmptyEntries
	}
	if payload.TestSplit < 0.0 || payload.TestSplit > 1.0 {
		return errorTrainSplitTestRange
	}

	var err error
	for _, entry := range payload.DataEntries {
		err = ValidateDataEntry(entry, inputsSize, labelsSize)
		if err != nil {
			return err
		}
	}
	return nil
}

// TrainPayloadFromJSON .
func TrainPayloadFromJSON(r io.Reader) (TrainPayload, error) {
	payload := TrainPayload{}
	err := json.NewDecoder(r).Decode(&payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
