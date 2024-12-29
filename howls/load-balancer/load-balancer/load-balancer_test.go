package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

func TestHandleTick(t *testing.T) {
	lb := loadBalancer{
		workers: make(map[string]worker),
		mutex:   sync.Mutex{},
		algorithm: map[string]algorithm{
			ROUND_ROBIN: newRoundRobin(),
		},
	}

	tests := []struct {
		name       string
		method     string
		payload    map[string]string
		wantStatus int
	}{
		{
			name:       "PUT method",
			method:     http.MethodPut,
			payload:    map[string]string{"address": "http://localhost:8081"},
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "DELETE method",
			method:     http.MethodDelete,
			payload:    map[string]string{"address": "http://localhost:8081"},
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "POST method",
			method:     http.MethodPost,
			payload:    map[string]string{"address": "http://localhost:8081"},
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "GET method",
			method:     http.MethodGet,
			payload:    nil,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var jsonPayload []byte
			if tt.payload != nil {
				jsonPayload, _ = json.Marshal(tt.payload)
			}
			req, err := http.NewRequest(tt.method, "/tick", bytes.NewBuffer(jsonPayload))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(lb.handlerTick)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}
		})
	}
}

func TestHandleStats(t *testing.T) {
	lb := loadBalancer{
		workers: make(map[string]worker),
		mutex:   sync.Mutex{},
		algorithm: map[string]algorithm{
			ROUND_ROBIN: newRoundRobin(),
		},
	}

	tests := []struct {
		name       string
		method     string
		wantStatus int
	}{
		{
			name:       "GET method with workers",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "GET method without workers",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "POST method",
			method:     http.MethodPost,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "PUT method",
			method:     http.MethodPut,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "DELETE method",
			method:     http.MethodDelete,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "GET method with workers" {
				lb.workers["http://localhost:8081"] = worker{}
			} else if tt.name == "GET method without workers" {
				lb.workers = make(map[string]worker)
			}

			req, err := http.NewRequest(tt.method, "/stats", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(lb.handlerStats)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}
		})
	}
}
func TestWorkerTick(t *testing.T) {
	lb := loadBalancer{
		workers: make(map[string]worker),
		mutex:   sync.Mutex{},
		algorithm: map[string]algorithm{
			ROUND_ROBIN: newRoundRobin(),
		},
	}

	tests := []struct {
		name       string
		address    string
		wantErr    bool
		wantWorker worker
	}{
		{
			name:    "Valid address",
			address: "http://localhost:8081",
			wantErr: false,
			wantWorker: worker{
				hostname: "localhost",
				url:      &url.URL{Scheme: "http", Host: "localhost:8081"},
				isAlive:  true,
				requests: 0,
			},
		},
		{
			name:    "Invalid address",
			address: "://invalid-url",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := lb.workerTick(tt.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("workerTick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				lb.mutex.Lock()
				defer lb.mutex.Unlock()
				gotWorker, exists := lb.workers[tt.wantWorker.hostname]
				if !exists {
					t.Errorf("workerTick() worker not registered")
					return
				}

				if gotWorker.hostname != tt.wantWorker.hostname || gotWorker.url.String() != tt.wantWorker.url.String() || gotWorker.isAlive != tt.wantWorker.isAlive || gotWorker.requests != tt.wantWorker.requests {
					t.Errorf("workerTick() gotWorker = %v, want %v", gotWorker, tt.wantWorker)
				}
			}
		})
	}
}
func TestWorkerRemove(t *testing.T) {
	lb := loadBalancer{
		workers: make(map[string]worker),
		mutex:   sync.Mutex{},
		algorithm: map[string]algorithm{
			ROUND_ROBIN: newRoundRobin(),
		},
	}

	tests := []struct {
		name     string
		hostname string
		setup    func()
		verify   func(t *testing.T)
	}{
		{
			name:     "Remove existing worker",
			hostname: "localhost",
			setup: func() {
				lb.workers["localhost"] = worker{
					hostname: "localhost",
					url:      &url.URL{Scheme: "http", Host: "localhost:8081"},
					isAlive:  true,
					requests: 0,
				}
			},
			verify: func(t *testing.T) {
				lb.mutex.Lock()
				defer lb.mutex.Unlock()
				if _, exists := lb.workers["localhost"]; exists {
					t.Error("worker should have been removed")
				}
			},
		},
		{
			name:     "Remove non-existing worker",
			hostname: "nonexistent",
			setup: func() {
				lb.workers = make(map[string]worker)
			},
			verify: func(t *testing.T) {
				lb.mutex.Lock()
				defer lb.mutex.Unlock()
				if len(lb.workers) != 0 {
					t.Error("workers map should be empty")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			lb.workerRemove(tt.hostname)
			tt.verify(t)
		})
	}
}
func TestUpdateWorkerStats(t *testing.T) {
	lb := loadBalancer{
		workers: make(map[string]worker),
		mutex:   sync.Mutex{},
		algorithm: map[string]algorithm{
			ROUND_ROBIN: newRoundRobin(),
		},
	}

	tests := []struct {
		name     string
		hostname string
		setup    func()
		verify   func(t *testing.T)
	}{
		{
			name:     "Update existing worker stats",
			hostname: "localhost",
			setup: func() {
				lb.workers["localhost"] = worker{
					hostname: "localhost",
					url:      &url.URL{Scheme: "http", Host: "localhost:8081"},
					isAlive:  true,
					requests: 5,
				}
			},
			verify: func(t *testing.T) {
				lb.mutex.Lock()
				defer lb.mutex.Unlock()
				if w, exists := lb.workers["localhost"]; exists {
					if w.requests != 6 {
						t.Errorf("requests count not incremented, got %d want %d", w.requests, 6)
					}
				} else {
					t.Error("worker should exist")
				}
			},
		},
		{
			name:     "Update non-existing worker stats",
			hostname: "nonexistent",
			setup: func() {
				lb.workers = make(map[string]worker)
			},
			verify: func(t *testing.T) {
				lb.mutex.Lock()
				defer lb.mutex.Unlock()
				if len(lb.workers) != 0 {
					t.Error("workers map should be empty")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			lb.updateWorkerStats(tt.hostname)
			tt.verify(t)
		})
	}
}