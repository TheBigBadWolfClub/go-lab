// nolint  funlen, gocognit,paralleltest
package reservation_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/reservation"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/client"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/test"
)

//nolint:gochecknoglobals
var (
	testHandler         reservation.HttpHandler
	testTableRepository table.Repository
	unlimitedTableID    int64
)

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Short() {
		db := test.DBOpen()
		defer db.Close()
		testTableRepository = table.NewRepository(db)
		testTableService := table.NewService(testTableRepository)

		testClientRepository := client.NewRepository(db)
		testClientService := client.NewService(testClientRepository, testTableService)

		clientListService := reservation.NewService(testClientService, testTableService)
		testHandler = reservation.NewHttpHandler(clientListService)

		// add scenario
		unlimitedTableID = test.SetupFindTableId(testTableService)
	}

	m.Run()
}

func Test_Handler_GET(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	tests := []struct {
		name       string
		method     string
		uri        string
		statusCode int
		expectBody bool
	}{
		{
			name:       "GET: expect StatusOK response",
			method:     http.MethodGet,
			uri:        table.URI,
			statusCode: http.StatusOK,
			expectBody: true,
		}, {
			name:       "GET: expect StatusNotFound response",
			method:     http.MethodPatch,
			uri:        table.URI,
			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// setup
			testTableRepository.Save(context.Background(), []int{7, 6})

			// prepare request
			req := httptest.NewRequest(tt.method, tt.uri, nil)
			w := httptest.NewRecorder()

			// invoke
			testHandler.Handler(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("GET status = %v, wantStatus %v", res.Status, tt.statusCode)
			}

			if tt.expectBody {
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("expected error to be nil got %v", err)
				}
				defer res.Body.Close()

				var payload reservation.RestModel
				json.Unmarshal(data, &payload)
				if len(data) == 0 {
					t.Errorf("expected body not to be nil: %v", data)
				}
			}
		})
	}
}

func Test_Handler_POST(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	clientName := fmt.Sprintf("ClientToReserve-%d", time.Now().UnixNano())
	tests := []struct {
		name            string
		method          string
		uri             string
		statusCode      int
		requestPayload  *reservation.RestModel
		responsePayload *reservation.RestModel
	}{
		{
			name:            "POST: expect StatusOK response",
			method:          http.MethodPost,
			uri:             reservation.URI + "/" + clientName,
			statusCode:      http.StatusOK,
			requestPayload:  &reservation.RestModel{GroupSize: 5, Table: unlimitedTableID},
			responsePayload: &reservation.RestModel{Name: clientName},
		}, {
			name:            "POST: expect StatusOK response",
			method:          http.MethodPost,
			uri:             reservation.URI + "/" + clientName,
			statusCode:      http.StatusConflict,
			requestPayload:  &reservation.RestModel{GroupSize: test.UnlimitedTableSeats + 1, Table: unlimitedTableID},
			responsePayload: &reservation.RestModel{Name: clientName},
		}, {
			name:       "POST: expect StatusUnprocessableEntity response",
			method:     http.MethodPost,
			uri:        reservation.URI + "/" + clientName,
			statusCode: http.StatusUnprocessableEntity,
		}, {
			name:       "POST: expect StatusNotFound, missing uri param \"name\"",
			method:     http.MethodPost,
			uri:        "/client",
			statusCode: http.StatusNotFound,
		}, {
			name:       "POST: expect StatusNotFound response",
			method:     http.MethodPatch,
			uri:        reservation.URI,
			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// prepare request
			w := httptest.NewRecorder()
			var req *http.Request
			if tt.requestPayload != nil {
				var buf bytes.Buffer
				_ = json.NewEncoder(&buf).Encode(tt.requestPayload)
				req = httptest.NewRequest(tt.method, tt.uri, &buf)
			} else {
				req = httptest.NewRequest(tt.method, tt.uri, nil)
			}

			// invoke
			testHandler.Handler(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("POST  status = %v, wantStatus %v", res.Status, tt.statusCode)
			}

			if tt.responsePayload != nil {
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("expected error to be nil got %v", err)
				}
				defer res.Body.Close()

				var payload reservation.RestModel
				json.Unmarshal(data, &payload)
				if reflect.DeepEqual(payload, tt.responsePayload) {
					t.Errorf("expected body not to be nil: %v", data)
				}
			}
		})
	}
}
