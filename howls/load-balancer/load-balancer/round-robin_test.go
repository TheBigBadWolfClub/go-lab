
package main

import (
	"testing"
)

func TestRoundRobin(t *testing.T) {
	type args struct {
		servers []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single server",
			args: args{
				servers: []string{"server1"},
			},
			want: "server1",
		},
		{
			name: "multiple servers",
			args: args{
				servers: []string{"server1", "server2", "server3"},
			},
			want: "server1",
		},
	}

	
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			rr := newRoundRobin()
			for _, server := range tt.args.servers {
				rr.registerWorker(server)
			}
			if got, _ := rr.getNextWorkerID(); got != tt.want {
				t.Errorf("roundRobin.getNextWorkerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundRobin_MultipleCalls(t *testing.T) {
	rr := newRoundRobin()
	servers := []string{"server1", "server2", "server3"}
	for _, server := range servers {
		rr.registerWorker(server)
	}

	expectedSequence := []string{"server1", "server2", "server3", "server1", "server2", "server3"}
	for i, expected := range expectedSequence {
		if got, _ := rr.getNextWorkerID(); got != expected {
			t.Errorf("roundRobin.getNextWorkerID() call %d = %v, want %v", i+1, got, expected)
		}
	}
}

func TestRoundRobin_NoServers(t *testing.T) {
	rr := newRoundRobin()
	if got, err := rr.getNextWorkerID(); err == nil || got != "" {
		t.Errorf("roundRobin.getNextWorkerID() with no servers = %v, %v, want empty string and error", got, err)
	}
}

func TestRoundRobin_RemoveWorker(t *testing.T) {
	rr := newRoundRobin()
	servers := []string{"server1", "server2", "server3"}
	for _, server := range servers {
		rr.registerWorker(server)
	}

	rr.removeWorker("server2")
	expectedSequence := []string{"server1", "server3", "server1", "server3"}
	for i, expected := range expectedSequence {
		if got, _ := rr.getNextWorkerID(); got != expected {
			t.Errorf("roundRobin.getNextWorkerID() call %d = %v, want %v", i+1, got, expected)
		}
	}
}

func TestRoundRobin_RegisterDuplicateWorker(t *testing.T) {
	rr := newRoundRobin()
	servers := []string{"server1", "server2", "server3"}
	for _, server := range servers {
		rr.registerWorker(server)
	}

	rr.registerWorker("server2") // Duplicate registration
	expectedSequence := []string{"server1", "server2", "server3", "server1", "server2", "server3"}
	for i, expected := range expectedSequence {
		if got, _ := rr.getNextWorkerID(); got != expected {
			t.Errorf("roundRobin.getNextWorkerID() call %d = %v, want %v", i+1, got, expected)
		}
	}
}