package main

import (
	"boston/learner"
	"boston/neural"
	"errors"
	"fmt"
	"sync"
)

var (
	errorNoSuchLearner        = errors.New("No Such Learner")
	errorLearnerAlreadyExists = errors.New("Learner Already Exists")
)

// LearnerMap .
type LearnerMap struct {
	sync.Mutex
	learners map[string]LearnerParams
}

// LearnerParams .
type LearnerParams struct {
	neural.NetworkConfig
	Mailbox chan learner.Action
}

// NewLearnerMap .
func NewLearnerMap() LearnerMap {
	return LearnerMap{
		learners: make(map[string]LearnerParams),
	}
}

// StartLearner puts a new learner in the map or
// if a learner with that
func (lmap *LearnerMap) StartLearner(name string, config neural.NetworkConfig) error {
	lmap.Lock()
	defer lmap.Unlock()
	_, ok := lmap.learners[name]
	if ok {
		return errorLearnerAlreadyExists
	}
	newLearner := learner.NewLearner(name, config)
	go newLearner.WaitForSignal()
	newLearnParams := LearnerParams{
		NetworkConfig: config,
		Mailbox:       newLearner.Mailbox,
	}
	lmap.learners[name] = newLearnParams
	return nil
}

// SendAction .
func (lmap *LearnerMap) SendAction(action learner.Action) error {
	name := action.LearnerName()
	lmap.Lock()
	defer lmap.Unlock()
	learnerParams, ok := lmap.learners[name]
	if !ok {
		return errorNoSuchLearner
	}
	if action.Signal() == learner.DeleteSignal {
		fmt.Println("deleting", action.LearnerName(), "from learners map")
		delete(lmap.learners, name)
	}
	learnerParams.Mailbox <- action
	return nil
}

// GetConfig returns the config for a
func (lmap *LearnerMap) GetConfig(name string) (neural.NetworkConfig, error) {
	lmap.Lock()
	defer lmap.Unlock()
	params, ok := lmap.learners[name]
	if !ok {
		return neural.NetworkConfig{}, errorNoSuchLearner
	}
	return params.NetworkConfig, nil
}

// Keys returns the current keys of learners
func (lmap *LearnerMap) Keys() []string {
	lmap.Lock()
	defer lmap.Unlock()
	keys := make([]string, 0, len(lmap.learners))
	for k := range lmap.learners {
		keys = append(keys, k)
	}
	return keys
}
