# Boston

A JSON HTTP server for configuring and using simple, concurrently-accessible neural networks.

Each neural network runs safely in it's own goroutine.

Heavily based on http://www.datadan.io/building-a-neural-net-from-scratch-in-go/


## Usage

1. To create a neural network `POST` the payload below to `http://localhost:4343/learners/create`:

	```json
	{
		"name": "example",
		"learning_rate": 0.1,
		"test_split": 0.3,
		"input_neurons": 10,
		"hidden_neurons": 4,
		"output_neurons": 4,
		"num_epochs": 5000
	}
	```

2. To train the neural network `POST` the payload below to `http://localhost:4343/learners/train`:

	```json
	{
		"name": "example",
		"entries": [
			{
				"inputs": [1.0, 0.0, 1.0, 1.0, 0.2, 1.0, 0.0, 1.0, 1.0, 0.2], 
				"labels": [0.0, 1.0, 0.1, 1.0]
			},
			{
				"inputs": [1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3],
				"labels": [0.0, 1.0, 0.1, 1.0]
			},
			{
				"inputs": [1.0, 0.0, 1.0, 1.0, 0.2, 1.0, 0.0, 1.0, 1.0, 0.2], 
				"labels": [0.0, 1.0, 0.1, 1.0]
			},
			{
				"inputs": [1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3],
				"labels": [0.0, 1.0, 0.1, 1.0]
			},
			{
				"inputs": [1.0, 0.0, 1.0, 1.0, 0.2, 1.0, 0.0, 1.0, 1.0, 0.2], 
				"labels": [0.0, 1.0, 0.1, 1.0]
			},
			{
				"inputs": [1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3],
				"labels": [0.0, 1.0, 0.1, 1.0]
			}
		],
		"test_split": 0.3
	}
	```

3. To predict inputs on a neural network `POST` the payload below to `http://localhost:4343/learners/predict`:

	```json
	{
		"name": "jason",
		"inputs": [
			[1.0, 0.0, 0.50, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3]
		]
	}
	```

4. To delete a neural network `POST` the payload below to `http://localhost:4343/learners/delete`:

	```json
	{
		"name": "jason"
	}
	```

5. To reset a neural network `POST` the payload below to `http://localhost:4343/learners/delete`:

	```json
	{
		"name": "jason"
	}
	```

# API 

#### Routes

```
/learners/list
/learners/create
/learners/delete
/learners/reset
/learners/train
/learners/predict
```

#### Full Example JSON:

```json
{
	"name": "example",
	"entries": [
		{
			"inputs": [1.0, 0.0, 1.0, 1.0, 0.2, 1.0, 0.0, 1.0, 1.0, 0.2], 
			"labels": [0.0, 1.0, 0.1, 1.0]
		},
		{
			"inputs": [1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3],
			"labels": [0.0, 1.0, 0.1, 1.0]
		},
		{
			"inputs": [1.0, 0.0, 1.0, 1.0, 0.2, 1.0, 0.0, 1.0, 1.0, 0.2], 
			"labels": [0.0, 1.0, 0.1, 1.0]
		},
		{
			"inputs": [1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3],
			"labels": [0.0, 1.0, 0.1, 1.0]
		},
		{
			"inputs": [1.0, 0.0, 1.0, 1.0, 0.2, 1.0, 0.0, 1.0, 1.0, 0.2], 
			"labels": [0.0, 1.0, 0.1, 1.0]
		},
		{
			"inputs": [1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3],
			"labels": [0.0, 1.0, 0.1, 1.0]
		}
	],
	"inputs": [
		[1.0, 0.0, 0.50, 1.0, 0.0, 0.98, 1.0, 0.0, 0.98, 0.3]
	],
	"learning_rate": 0.1,
	"test_split": 0.3,
	"input_neurons": 10,
	"hidden_neurons": 4,
	"output_neurons": 4,
	"num_epochs": 5000
}
```