package main

import (
	"fmt"
	"sync"
)

type roundRobin struct {
	order   []string
	current int
	mutex   sync.Mutex
}

func newRoundRobin() *roundRobin {
	return &roundRobin{
		order:   []string{},
		current: 0,
		mutex:   sync.Mutex{},
	}
}

func (rb *roundRobin) getNextWorkerID() (string, error) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()
	if len(rb.order) == 0 {
		return "", fmt.Errorf("no workers available")
	}

	if rb.current >= len(rb.order) {
		rb.current = 0
	}

	host := rb.order[rb.current]
	rb.current++
	return host, nil
}

func (rb *roundRobin) registerWorker(address string) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()
	for _, host := range rb.order {
		if host == address {
			return
		}
	}

	rb.order = append(rb.order, address)
}

func (rb *roundRobin) removeWorker(address string) {
	rb.mutex.Lock()
	defer rb.mutex.Unlock()
	for i, host := range rb.order {
		if host == address {
			rb.order = append(rb.order[:i], rb.order[i+1:]...)
			break
		}
	}
}
