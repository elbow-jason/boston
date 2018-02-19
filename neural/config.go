package neural

// NetworkConfig defines our neural network
// architecture and learning parameters.
type NetworkConfig struct {
	InputNeurons  int     `json:"input_neurons"`
	HiddenNeurons int     `json:"hidden_neurons"`
	OutputNeurons int     `json:"output_neurons"`
	NumEpochs     int     `json:"num_epochs"`
	LearningRate  float64 `json:"learning_rate"`
}
