package main

import (
	"errors"
	"sync"
)

var (
	errorNoSuchNetwork = errors.New("No Such Network")
)

// NetworkMap .
type NetworkMap struct {
	sync.Mutex
	networks map[string]*NeuralNet
}

// PutNetwork .
func (nmap *NetworkMap) PutNetwork(name string, network *NeuralNet) {
	nmap.Lock()
	nmap.networks[name] = network
	nmap.Unlock()
}

// GetNetwork .
func (nmap *NetworkMap) GetNetwork(name string) (*NeuralNet, error) {
	nmap.Lock()
	nn, ok := nmap.networks[name]
	nmap.Unlock()
	if !ok {
		return nil, errorNoSuchNetwork
	}
	return nn, nil
}

// DeleteNetwork .
func (nmap *NetworkMap) DeleteNetwork(name string) {
	nmap.Lock()
	defer nmap.Unlock()
	_, ok := nmap.networks[name]
	if ok {
		delete(nmap.networks, name)
	}

}
