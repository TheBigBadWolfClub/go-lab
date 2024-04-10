package presentation

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// https://newrelic.com/blog/best-practices/opentelemetry-histograms

type GeneratorAPI interface {
	GenerateDiceRollers(http.ResponseWriter, *http.Request)
}

type generatorApi struct {
}

func NewGeneratorAPI() GeneratorAPI {
	return &generatorApi{}
}

func (a *generatorApi) GenerateDiceRollers(w http.ResponseWriter, r *http.Request) {

	nRollsStr := r.URL.Query().Get("number_of_rolls")

	if nRollsStr == "" {
		nRollsStr = "1"
	}

	nRolls, err := strconv.Atoi(nRollsStr)
	if err != nil {
		log.Printf("Error converting string to int: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respStr := fmt.Sprintf("Number of Rolls: %d\n", nRolls)
	mapResp := make(map[int]int)

	results := make(chan int, nRolls)
	for i := 0; i < nRolls; i++ {
		go func() {
			resp, err := http.Get("http://localhost:8091/api/v1/rolldice")
			if err != nil {
				results <- 505
				return
			}

			data, _ := io.ReadAll(resp.Body)
			strData := string(data)
			intValue, err := strconv.Atoi(strings.ReplaceAll(strData, "\n", ""))
			if err != nil {
				results <- 505
				return
			}

			results <- intValue
		}()
	}

	for i := 0; i < nRolls; i++ {
		select {
		case result := <-results:
			mapResp[i] = result
		}
	}

	mapRespStr := fmt.Sprintf("%v", mapResp)
	_, _ = w.Write([]byte(respStr + mapRespStr))
}
