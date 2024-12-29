package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

const ROUND_ROBIN = "round-robin"

type loadBalancer struct {
	workers   map[string]worker
	mutex     sync.Mutex
	algorithm map[string]algorithm
}


type worker struct {
	hostname string
	isAlive  bool
	requests int
	url      *url.URL
}

type algorithm interface {
	getNextWorkerID() (string, error)
	registerWorker(address string)
	removeWorker(address string)
}

func main() {

	lb := loadBalancer{
		workers: make(map[string]worker),
		mutex:   sync.Mutex{},
		algorithm: map[string]algorithm{
			ROUND_ROBIN: newRoundRobin(),
		},
	}

	http.HandleFunc("/tick", lb.handlerTick)
	http.HandleFunc("/stats", lb.handlerStats)
	http.HandleFunc("/", lb.proxyRequest)

	// Default handler for non-registered routes
	http.HandleFunc("/not-found", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}

}

func (lb *loadBalancer) handlerTick(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodDelete {
		http.Error(w, "404 - Not Found", http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	host := struct {
		Address string `json:"address"`
	}{}
	err = json.Unmarshal(body, &host)
	if err != nil {
		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPut {
		err := lb.workerTick(host.Address)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("fail to add host"))
			return
		}

		w.WriteHeader(http.StatusNoContent) // 204 No Content
		return
	}

	if r.Method == http.MethodDelete {
		lb.workerRemove(host.Address)
		w.WriteHeader(http.StatusNoContent) // 204 No Content
		return
	}
}

func (lb *loadBalancer) handlerStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "404 - Not Found", http.StatusNotFound)
		return
	}
	
	type wStats struct {
		Worker string `json:"worker"`
		IsLive bool   `json:"is_alive"`
		NReq   int    `json:"requests"`
	}
	stats := make([]wStats, 0, len(lb.workers))

	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	for _, srv := range lb.workers {
		stats = append(stats, wStats{
			Worker: srv.hostname,
			IsLive: srv.isAlive,
			NReq:   srv.requests,
		})
	}

	statsJson, err := json.Marshal(stats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("fail to marshal stats %v\n", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(statsJson)
}

func (lb *loadBalancer) proxyRequest(w http.ResponseWriter, r *http.Request) {
	wrk, err := lb.getWorker(ROUND_ROBIN)
	if err != nil {
		http.Error(w, "No workers available", http.StatusServiceUnavailable)
		return
	}

	lb.processRequest(w, r, wrk)
	lb.updateWorkerStats(wrk.hostname)
}

func (lb *loadBalancer) workerTick(address string) error {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	//hostUrl, err := url.Parse("http://127.0.0.1:8080")
	hostUrl, err := url.Parse(address)
	if err != nil {
		return err
	}

	srv, ok := lb.workers[hostUrl.Hostname()]
	if !ok {
		srv = worker{
			hostname: hostUrl.Hostname(),
			url:      hostUrl,
			isAlive:  true,
			requests: 0,
		}

	}

	srv.isAlive = true
	lb.workers[hostUrl.Hostname()] = srv

	lb.algorithm[ROUND_ROBIN].registerWorker(address)
	return nil
}

func (lb *loadBalancer) workerRemove(hostname string) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	for i, srv := range lb.workers {
		if srv.hostname == hostname {
			delete(lb.workers, i)
			break
		}
	}

	lb.algorithm[ROUND_ROBIN].removeWorker(hostname)
}

func (lb *loadBalancer) getWorker(algorithm string) (worker, error) {
	if _, ok := lb.algorithm[algorithm]; !ok {
		return worker{}, fmt.Errorf("algorithm not found")
	}

	lb.mutex.Lock()
	defer lb.mutex.Unlock()
	for len(lb.workers) > 0 {
		id, err := lb.algorithm[algorithm].getNextWorkerID()
		if err != nil {
			return worker{}, err
		}

		if w, ok := lb.workers[id]; ok {
			if w.isAlive {
				return w, nil
			}
		}

		lb.workerRemove(id)
	}

	return worker{}, fmt.Errorf("no workers available")
}

func (lb *loadBalancer) processRequest(w http.ResponseWriter, r *http.Request, wrk worker) {
	w.Header().Add("X-Forwarded-Server", wrk.url.String())
	proxy := httputil.NewSingleHostReverseProxy(wrk.url)
	proxy.ServeHTTP(w, r)

	fmt.Printf("Request proxied to %s\n", wrk.url.String())
}

func (lb *loadBalancer) updateWorkerStats(hostname string) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	w, ok := lb.workers[hostname]
	if ok {
		w.requests++
		lb.workers[hostname] = w
	}

}
