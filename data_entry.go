package main

import "errors"

var (
	errorEmptyLabels = errors.New("Entry 'labels' cannot be empty")
	errorEmptyInputs = errors.New("Entry 'inputs' cannot be empty")
	errorLabelsSize  = errors.New("All entry 'labels' must be a uniform size")
	errorInputsSize  = errors.New("All entry 'inputs' must be a uniform size")
)

// DataEntry .
type DataEntry struct {
	RawInputs []float64 `json:"inputs"`
	RawLabels []float64 `json:"labels"`
}

// Validate ensures an entry does not contain empty values
func (d DataEntry) Validate(expectedInputsSize, expectedLabelsSize int) error {
	inputsSize := len(d.RawInputs)
	labelsSize := len(d.RawLabels)
	if inputsSize == 0 {
		return errorEmptyInputs
	}
	if labelsSize == 0 {
		return errorEmptyLabels
	}
	if inputsSize != expectedInputsSize {
		return errorInputsSize
	}
	if labelsSize != expectedLabelsSize {
		return errorLabelsSize
	}
	return nil
}
