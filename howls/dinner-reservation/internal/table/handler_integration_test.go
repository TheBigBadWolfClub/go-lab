// nolint funlen, gocognit,paralleltest
package table_test

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/test"
)

//nolint:gochecknoglobals
var (
	testRepository table.Repository
	testHandler    table.HttpHandler
)

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Short() {
		db := test.DBOpen()
		defer db.Close()
		testRepository = table.NewRepository(db)
		testService := table.NewService(testRepository)
		testHandler = table.NewHttpHandler(testService)
	}

	m.Run()
}

func Test_Handler_GET(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	tests := []struct {
		name            string
		method          string
		uri             string
		statusCode      int
		responsePayload bool
	}{
		{
			name:            "GET: expect StatusOK response",
			method:          http.MethodGet,
			uri:             table.URI,
			statusCode:      http.StatusOK,
			responsePayload: true,
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
			testRepository.Save(context.Background(), []int{4, 5})

			// prepare request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, tt.uri, nil)

			// invoke
			testHandler.Handler(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("GET status = %v, wantStatus %v", res.Status, tt.statusCode)
			}

			if tt.responsePayload {
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("expected error to be nil got %v", err)
				}
				defer res.Body.Close()

				var payload []table.RestModel
				json.Unmarshal(data, &payload)
				if len(payload) == 0 {
					t.Errorf("expected body not to be nil: %v", string(data))
				}
			}
		})
	}
}

func Test_Handler_DELETE(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	tests := []struct {
		name       string
		method     string
		uri        string
		statusCode int
	}{
		{
			name:       "DELETE: expect StatusOK response",
			method:     http.MethodDelete,
			uri:        table.URI,
			statusCode: http.StatusNoContent,
		}, {
			name:   "DELETE: expect StatusNotFound response",
			method: http.MethodPatch,
			uri:    table.URI,

			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// setup
			testRepository.Delete(context.Background())

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

func Test_Handler_POST(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests!")
	}

	tests := []struct {
		name       string
		method     string
		uri        string
		payload    *[]int
		statusCode int
	}{
		{
			name:       "POST: expect StatusOK response",
			method:     http.MethodPost,
			uri:        table.URI,
			statusCode: http.StatusNoContent,
			payload:    &[]int{1, 2, 3},
		},
		{
			name:       "POST: expect StatusUnprocessableEntity response",
			method:     http.MethodPost,
			uri:        table.URI,
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:       "POST: expect StatusNotFound response",
			method:     http.MethodPatch,
			uri:        table.URI,
			statusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// prepare request
			w := httptest.NewRecorder()
			var req *http.Request
			if tt.payload != nil {
				var buf bytes.Buffer
				_ = json.NewEncoder(&buf).Encode(tt.payload)
				req = httptest.NewRequest(tt.method, tt.uri, &buf)
			} else {
				req = httptest.NewRequest(tt.method, tt.uri, nil)
			}

			// invoke
			testHandler.Handler(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.statusCode {
				t.Errorf("POST status = %v, wantStatus %v", res.Status, tt.statusCode)
			}
		})
	}
}
