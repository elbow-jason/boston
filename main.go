package main

import (
	"boston/learner"
	"boston/neural"
	"boston/server"
	"log"
)

func main() {

	action, err := learner.NewCreatePayload("jason", 10, 4, 4, 5000, 0.1)
	if err != nil {
		log.Fatal(err)
	}
	learnerMap := learner.NewMap()
	config := neural.NetworkConfig{
		InputNeurons:  action.InputNeurons,
		HiddenNeurons: action.HiddenNeurons,
		OutputNeurons: action.OutputNeurons,
		LearningRate:  action.LearningRate,
		NumEpochs:     action.NumEpochs,
	}
	learnerMap.StartLearner("jason", config)
	// learnerMap.SendAction("jason", action)
	// runPrediction()
	server.StartLearnerHTTPServer(&learnerMap)
}

// func runPrediction() {
// 	config := NeuralNetConfig{
// 		InputNeurons:  4,
// 		OutputNeurons: 3,
// 		HiddenNeurons: 3,
// 		NumEpochs:     5000,
// 		LearningRate:  0.3,
// 	}
// 	// Form the training matrices.
// 	inputs, labels := makeInputsAndLabels("data/iris_train.csv")

// 	// Define our network architecture and learning parameters.

// 	// Train the neural network.
// 	network := NewNeuralNetwork(config)

// 	if err := network.train(inputs, labels); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Form the testing matrices.
// 	testInputs, testLabels := makeInputsAndLabels("data/iris_test.csv")

// 	// Make the predictions using the trained model.
// 	predictions, err := network.predict(testInputs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Calculate the accuracy of our model.
// 	var truePosNeg int
// 	numPreds, _ := predictions.Dims()
// 	for i := 0; i < numPreds; i++ {

// 		// Get the label.
// 		labelRow := mat.Row(nil, i, testLabels)
// 		var prediction int
// 		for idx, label := range labelRow {
// 			if label == 1.0 {
// 				prediction = idx
// 				break
// 			}
// 		}

// 		// Accumulate the true positive/negative count.
// 		if predictions.At(i, prediction) == floats.Max(mat.Row(nil, i, predictions)) {
// 			truePosNeg++
// 		}
// 	}

// 	// Calculate the accuracy (subset accuracy).
// 	accuracy := float64(truePosNeg) / float64(numPreds)

// 	// Output the Accuracy value to standard out.
// 	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
// }
