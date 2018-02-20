package main

import (
	"boston/learner"
	"boston/server"
)

// action, err := learner.NewCreatePayload("jason", 10, 4, 4, 5000, 0.1)
// if err != nil {
// 	log.Fatal(err)
// }
//
// config := neural.NetworkConfig{
// 	InputNeurons:  action.InputNeurons,
// 	HiddenNeurons: action.HiddenNeurons,
// 	OutputNeurons: action.OutputNeurons,
// 	LearningRate:  action.LearningRate,
// 	NumEpochs:     action.NumEpochs,
// }
// learnerMap.StartLearner("jason", config)

func main() {
	learnerMap := learner.NewMap()
	server.StartLearnerHTTPServer(&learnerMap, ":4343")
}
