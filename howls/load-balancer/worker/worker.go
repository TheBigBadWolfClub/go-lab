package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"time"

	"os"
)

const DEFAULT_PORT = "8080"
const DEFAULT_ADDRESS = "localhost"
const DEFAULT_LOAD_BALANCER = "http://localhost:80"
const DEFAULT_REGISTER_RETRY = "1s"
const DEFAULT_TICK_RETRY = "6s"

const ENV_PORT = "PORT"
const ENV_LOAD_BALANCER_ADDRESS = "LOAD_BALANCER_ADDRESS"
const ENV_REGISTER_RETRY = "REGISTER_RETRY"
const ENV_TICK_RETRY = "TICK_RETRY"

type notRegisterError error

type worker struct {
	address             string
	loadBalancerAddress string
	registerRetry       time.Duration
	tickRetry           time.Duration
}

func main() {

	wrk := newWorker()
	fmt.Printf("worker: %+v\n", wrk)

	go wrk.deployProbe()
	defer wrk.safeUnregister()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		src := rand.NewSource(time.Now().UnixNano())
		intRand := rand.New(src)
		randomInt := intRand.Intn(2000)
		time.Sleep(time.Duration(randomInt) * time.Millisecond)

		reply := fmt.Sprintf("Worker: %s, Work ms: %d\n", wrk.address, randomInt)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(reply))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func newWorker() *worker {

	// WORKER
	port := os.Getenv(ENV_PORT)
	if port == "" {
		port = DEFAULT_PORT
	}

	address, err := getIpAddress()
	if err != nil {
		address = DEFAULT_ADDRESS
	}

	// LOAD BALANCER
	lbAddress := os.Getenv(ENV_LOAD_BALANCER_ADDRESS)
	if lbAddress == "" {
		lbAddress = DEFAULT_LOAD_BALANCER
	}

	// TICK RETRY
	ticketRetryStr := os.Getenv(ENV_TICK_RETRY)
	if ticketRetryStr == "" {
		ticketRetryStr = DEFAULT_TICK_RETRY
	}

	ticketRetry, err := time.ParseDuration(ticketRetryStr)
	if err != nil {
		panic(fmt.Sprintf("error parsing TICK RETRY: %v", err))
	}

	// REGISTER RETRY
	registerRetryStr := os.Getenv(ENV_REGISTER_RETRY)
	if registerRetryStr == "" {
		registerRetryStr = DEFAULT_REGISTER_RETRY
	}
	registerRetry, err := time.ParseDuration(registerRetryStr)
	if err != nil {
		panic(fmt.Sprintf("error parsing REGISTER RETRY: %v", err))
	}

	wrk := worker{
		address:             address + ":" + port,
		loadBalancerAddress: lbAddress,
		tickRetry:           ticketRetry,
		registerRetry:       registerRetry,
	}

	fmt.Printf("worker config: %+v\n", wrk)
	return &wrk
}

func getIpAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("error getting interface addresses: %v", err)
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				fmt.Printf("ip address: %s\n", ipNet.IP.String())
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("fail to get ip address")
}

func (w *worker) deployProbe() {
registerWork:
	var notError notRegisterError
	err := w.registerWorker()
	if err != nil && !errors.As(err, &notError) {
		panic(fmt.Sprintf("error registering worker: %v", err))
	}
	goto tickWorker

tickWorker:
	err = w.tickProbe()
	if err != nil {
		goto registerWork
	}

	goto tickWorker
}

func (w *worker) tickProbe() error {

	request, err := w.buildRegisterRequest()
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	ticker := time.NewTicker(w.tickRetry)
	for range ticker.C {
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Printf("error request tick: %v\n", err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode >= http.StatusBadRequest {
			return fmt.Errorf("fail to tick: %s", w.address)
		}

		fmt.Printf("tick probe is alive: %s\n", w.address)
	}
	return nil
}

func (w *worker) registerWorker() error {
	request, err := w.buildRegisterRequest()
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	ticker := time.NewTicker(w.registerRetry)
	for range ticker.C {
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Printf("error making request, registering worker: %v\n", err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode >= http.StatusBadRequest {
			fmt.Printf("error registering worker, status code: %d\n", resp.StatusCode)
			return notRegisterError(fmt.Errorf("%d", resp.StatusCode))
		}

		fmt.Printf("worker registered: %s\n", w.address)
		break
	}
	return nil
}

func (w *worker) unregisterWorker() error {
	request, err := w.buildRegisterRequest()
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to unregister worker, status code: %d", resp.StatusCode)
	}
	return nil
}

func (w *worker) safeUnregister() {
	var err error
	if r := recover(); r != nil {
		err = w.unregisterWorker()
	}
	err = w.unregisterWorker()

	if err != nil {
		fmt.Printf("error unregistering worker: %v", err)
	}
}

func (w *worker) buildRegisterRequest() (*http.Request, error) {
	payload := map[string]string{
		"address": fmt.Sprintf("http://%s", w.address),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	req, err := http.NewRequest("POST", w.loadBalancerAddress+"/tick", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	return req, nil
}
