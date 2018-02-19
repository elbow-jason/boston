package learner

import (
	// "boston/learner"
	// "boston/learner"
	"boston/neural"
	"errors"
	"fmt"
	"sync"
)

var (
	errorNoSuchLearner        = errors.New("No Such Learner")
	errorLearnerAlreadyExists = errors.New("Learner Already Exists")
)

// Map .
type Map struct {
	sync.Mutex
	learners map[string]Params
}

// Params .
type Params struct {
	neural.NetworkConfig
	Mailbox chan Action
}

// NewMap .
func NewMap() Map {
	return Map{
		learners: make(map[string]Params),
	}
}

// StartLearner puts a new learner in the map or
// if a learner with that
func (lmap *Map) StartLearner(name string, config neural.NetworkConfig) error {
	lmap.Lock()
	defer lmap.Unlock()
	_, ok := lmap.learners[name]
	if ok {
		return errorLearnerAlreadyExists
	}
	newLearner := NewLearner(name, config)
	go newLearner.WaitForSignal()
	newParams := Params{
		NetworkConfig: config,
		Mailbox:       newLearner.Mailbox,
	}
	lmap.learners[name] = newParams
	return nil
}

// SendAction .
func (lmap *Map) SendAction(action Action) error {
	name := action.Name()
	lmap.Lock()
	defer lmap.Unlock()
	learnerParams, ok := lmap.learners[name]
	if !ok {
		return errorNoSuchLearner
	}
	if action.Signal() == DeleteSignal {
		fmt.Println("deleting", action.Name(), "from learners map")
		delete(lmap.learners, name)
	}
	learnerParams.Mailbox <- action
	return nil
}

// GetConfig returns the config for a
func (lmap *Map) GetConfig(name string) (neural.NetworkConfig, error) {
	lmap.Lock()
	defer lmap.Unlock()
	params, ok := lmap.learners[name]
	if !ok {
		return neural.NetworkConfig{}, errorNoSuchLearner
	}
	return params.NetworkConfig, nil
}

// Keys returns the current keys of learners
func (lmap *Map) Keys() []string {
	lmap.Lock()
	defer lmap.Unlock()
	keys := make([]string, 0, len(lmap.learners))
	for k := range lmap.learners {
		keys = append(keys, k)
	}
	return keys
}
