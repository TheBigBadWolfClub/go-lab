// nolint  funlen, gocognit,paralleltest,gochecknoglobals
package client_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/client"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/test"
)

var (
	testRepository   client.Repository
	testHandler      client.HttpHandler
	unlimitedTableID int64
)

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Short() {
		db := test.DBOpen()
		defer db.Close()

		// dependency
		tableService := table.NewService(table.NewRepository(db))

		testRepository = client.NewRepository(db)
		testService := client.NewService(testRepository, tableService)
		testHandler = client.NewHttpHandler(testService)

		// add scenario
		unlimitedTableID = test.SetupFindTableId(tableService)
	}

	m.Run()
}

func Test_Handler_GET_ClientsHavingDinner(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	// setup
	testRepository.Save(context.Background(), client.Client{
		Name:    "ClientCheckIn",
		Size:    2,
		TableID: unlimitedTableID,
	})

	testRepository.UpdateCheckIn(context.Background(), client.Client{
		Name:    "ClientCheckIn",
		Size:    3,
		TableID: unlimitedTableID,
	})

	// prepare request
	req := httptest.NewRequest(http.MethodGet, client.URI, nil)
	w := httptest.NewRecorder()

	// invoke
	testHandler.Handler(w, req)

	// validate result
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("GET status = %v, wantStatus %v", res.Status, http.StatusOK)
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if !strings.Contains(string(data), "ClientCheckIn") {
		t.Errorf("client not in party: %v", string(data))
	}
}

func Test_Handler_PUT_CheckIn(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	// setup
	clientName := "ClientToCheckIn"
	testRepository.Save(context.Background(), client.Client{
		Name:    clientName,
		Size:    5,
		TableID: unlimitedTableID,
	})

	tests := []struct {
		name            string
		method          string
		uri             string
		statusCode      int
		requestPayload  *client.RestModel
		responsePayload *client.RestModel
	}{
		{
			name:            "PUT: expect StatusOK response",
			method:          http.MethodPut,
			uri:             client.URI + "/" + clientName,
			statusCode:      http.StatusOK,
			requestPayload:  &client.RestModel{GroupSize: 5},
			responsePayload: &client.RestModel{Name: clientName},
		}, {
			name:            "PUT: expect StatusConflict response, to many clients",
			method:          http.MethodPut,
			uri:             client.URI + "/" + clientName,
			statusCode:      http.StatusConflict,
			requestPayload:  &client.RestModel{GroupSize: test.UnlimitedTableSeats + 1},
			responsePayload: &client.RestModel{Name: clientName},
		}, {
			name:       "PUT: expect StatusUnprocessableEntity response",
			method:     http.MethodPut,
			uri:        client.URI + "/clientWillArrive",
			statusCode: http.StatusUnprocessableEntity,
		}, {
			name:       "PUT: expect StatusNotFound, missing uri param \"name\"",
			method:     http.MethodPut,
			uri:        "/client",
			statusCode: http.StatusNotFound,
		}, {
			name:       "PUT: expect StatusNotFound response",
			method:     http.MethodPatch,
			uri:        client.URI,
			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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
				t.Errorf("PUT  status = %v, wantStatus %v", res.Status, tt.statusCode)
			}

			if tt.responsePayload != nil {
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("expected error to be nil got %v", err)
				}
				defer res.Body.Close()

				var payload client.RestModel
				json.Unmarshal(data, &payload)
				if reflect.DeepEqual(payload, tt.responsePayload) {
					t.Errorf("expected body not to be nil: %v", data)
				}
			}
		})
	}
}

func Test_Handler_DELETE(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	// setup
	clientName := "ClientToCheckOut"
	testRepository.Save(context.Background(), client.Client{
		Name:    clientName,
		Size:    5,
		TableID: unlimitedTableID,
	})

	tests := []struct {
		name       string
		method     string
		uri        string
		statusCode int
	}{
		{
			name:       "DELETE: expect StatusNoContent response",
			method:     http.MethodDelete,
			uri:        client.URI + "/" + clientName,
			statusCode: http.StatusNoContent,
		}, {
			name:       "DELETE: expect StatusNotFound, missing uri param \"name\"",
			method:     http.MethodDelete,
			uri:        client.URI,
			statusCode: http.StatusNotFound,
		}, {
			name:       "DELETE: expect StatusNotFound response",
			method:     http.MethodPatch,
			uri:        client.URI,
			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// prepare request
			req := httptest.NewRequest(tt.method, tt.uri, nil)
			w := httptest.NewRecorder()

			// invoke
			testHandler.Handler(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("DELETE status = %v, wantStatus %v", res.Status, tt.statusCode)
			}
		})
	}
}
