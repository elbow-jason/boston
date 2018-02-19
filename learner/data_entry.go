package learner

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

var (
	errorEmptyLabels  = errors.New("Entry 'labels' cannot be empty")
	errorEmptyInputs  = errors.New("Entry 'inputs' cannot be empty")
	errorEmptyEntries = errors.New("Entry 'entries' cannot be empty")
)

func makeErrorLabelsSize(size int) error {
	message := fmt.Sprintf("All entry 'labels' sizes must be %d ('output_neurons')", size)
	return errors.New(message)
}

func makeErrorInputsSize(size int) error {
	message := fmt.Sprintf("All entry 'inputs' sizes must be %d ('input_neurons')", size)
	return errors.New(message)
}

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
		return makeErrorInputsSize(expectedInputsSize)
	}
	if labelsSize != expectedLabelsSize {
		return makeErrorLabelsSize(expectedLabelsSize)
	}
	return nil
}

// DataEntriesToMatrices turns entries into
// inputs and labels matrices
func DataEntriesToMatrices(entries []DataEntry) (*mat.Dense, *mat.Dense) {
	rows := len(entries)
	firstRow := entries[0]
	inputs := []float64{}
	inputCols := len(firstRow.RawInputs)
	labels := []float64{}
	labelsCols := len(firstRow.RawLabels)
	for _, entry := range entries {
		inputs = append(inputs, entry.RawInputs...)
		labels = append(labels, entry.RawLabels...)
	}
	return mat.NewDense(rows, inputCols, inputs), mat.NewDense(rows, labelsCols, labels)
}
