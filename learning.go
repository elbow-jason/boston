package main

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"errors"
// 	"io"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"time"

// 	"gonum.org/v1/gonum/mat"
// )

// var (
// 	errorNameRequired      = errors.New("Learning job 'name' is required")
// 	errorEmptyEntries      = errors.New("Learning job 'entries' cannot be empty")
// 	errorTestSplitRange    = errors.New("Learning job 'test_split' must be between 0.0 and 1.0")
// 	errorLearningRateRange = errors.New("Learning job 'learning_rate' must be between 0.0 and 1.0")
// 	// errorInputNeuronsRequired  = errors.New("Learning job 'input_neurons' must be greater than 0")
// 	// errorHiddenNeuronsRequired = errors.New("Learning job 'hidden_neurons' must be greater than 0")
// 	// errorOutputNeuronsRequired = errors.New("Learning job 'output_neurons' must be greater than 0")
// 	// errorNumEpochsRequired     = errors.New("Learning job 'num_epochs' must be greater than 0")
// )

// // LearningJob .
// type LearningJob struct {
// 	Name          string      `json:"name"`
// 	Entries       []DataEntry `json:"entries"`
// 	TestSplit     float64     `json:"test_split"`
// 	LearningRate  float64     `json:"learning_rate"`
// 	InputNeurons  int         `json:"input_neurons"`
// 	HiddenNeurons int         `json:"hidden_neurons"`
// 	OutputNeurons int         `json:"output_neurons"`
// 	NumEpochs     int         `json:"num_epochs"`
// }

// // ToNeuralNetConfig .
// func (l LearningJob) ToNeuralNetConfig() NeuralNetConfig {
// 	return NeuralNetConfig{
// 		LearningRate:  l.LearningRate,
// 		InputNeurons:  l.InputNeurons,
// 		HiddenNeurons: l.HiddenNeurons,
// 		OutputNeurons: l.OutputNeurons,
// 		NumEpochs:     l.NumEpochs,
// 	}
// }

// // TestAndTrainingSets .
// func (l LearningJob) TestAndTrainingSets() (trainEntries []DataEntry, testEntries []DataEntry) {
// 	randSource := rand.NewSource(time.Now().UnixNano())
// 	randGenerator := rand.New(randSource)
// 	testSplit := l.TestSplit
// 	for _, entry := range l.Entries {
// 		randNum := randGenerator.Float64()
// 		if randNum <= testSplit {
// 			testEntries = append(testEntries, entry)
// 		} else {
// 			trainEntries = append(trainEntries, entry)
// 		}

// 	}
// 	return
// }

// // Validate ensures a learning job has entries and is properly configured.
// func (l LearningJob) Validate() error {
// 	if l.Name == "" {
// 		return errorNameRequired
// 	}
// 	if len(l.Entries) == 0 {
// 		return errorEmptyEntries
// 	}
// 	if l.TestSplit <= 0.0 || l.TestSplit > 1.0 {
// 		return errorTestSplitRange
// 	}
// 	if l.LearningRate <= 0.0 || l.LearningRate > 1.0 {
// 		return errorLearningRateRange
// 	}
// 	// if l.InputNeurons == 0 {
// 	// 	return errorInputNeuronsRequired
// 	// }
// 	// if l.HiddenNeurons == 0 {
// 	// 	return errorHiddenNeuronsRequired
// 	// }
// 	// if l.OutputNeurons == 0 {
// 	// 	return errorOutputNeuronsRequired
// 	// }
// 	// if l.NumEpochs == 0 {
// 	// 	return errorNumEpochsRequired
// 	// }
// 	firstEntry := l.Entries[0]
// 	inputsSize := len(firstEntry.RawInputs)
// 	labelsSize := len(firstEntry.RawLabels)
// 	for _, entry := range l.Entries {
// 		if err := entry.Validate(inputsSize, labelsSize); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // InputsMatrix .
// func (l LearningJob) InputsMatrix() *mat.Dense {
// 	var inputsData []float64
// 	rows := len(l.Entries)

// 	return mat.NewDense(rows, 4, inputsData)
// }

// func learingJobFromJSON(jsonBytes io.Reader) (LearningJob, error) {
// 	job := LearningJob{}
// 	err := json.NewDecoder(jsonBytes).Decode(&job)
// 	if err != nil {
// 		return job, err
// 	}
// 	if err = job.Validate(); err != nil {
// 		return job, err
// 	}
// 	return job, nil
// }

// func makeInputsAndLabels(fileName string) (*mat.Dense, *mat.Dense) {
// 	// Open the dataset file.
// 	f, err := os.Open(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	// Create a new CSV reader reading from the opened file.
// 	reader := csv.NewReader(f)
// 	reader.FieldsPerRecord = 7

// 	// Read in all of the CSV records
// 	rawCSVData, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// inputsData and labelsData will hold all the
// 	// float values that will eventually be
// 	// used to form matrices.
// 	inputsData := make([]float64, 4*len(rawCSVData))
// 	labelsData := make([]float64, 3*len(rawCSVData))

// 	// Will track the current index of matrix values.
// 	var inputsIndex int
// 	var labelsIndex int

// 	// Sequentially move the rows into a slice of floats.
// 	for idx, record := range rawCSVData {

// 		// Skip the header row.
// 		if idx == 0 {
// 			continue
// 		}

// 		// Loop over the float columns.
// 		for i, val := range record {

// 			// Convert the value to a float.
// 			parsedVal, err := strconv.ParseFloat(val, 64)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			// Add to the labelsData if relevant.
// 			if i == 4 || i == 5 || i == 6 {
// 				labelsData[labelsIndex] = parsedVal
// 				labelsIndex++
// 				continue
// 			}

// 			// Add the float value to the slice of floats.
// 			inputsData[inputsIndex] = parsedVal
// 			inputsIndex++
// 		}
// 	}
// 	dataSetSize := len(rawCSVData)
// 	inputs := mat.NewDense(dataSetSize, 4, inputsData)
// 	labels := mat.NewDense(dataSetSize, 3, labelsData)
// 	return inputs, labels
// }
