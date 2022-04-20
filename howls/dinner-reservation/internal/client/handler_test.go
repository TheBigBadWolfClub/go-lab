// nolint testpackage, funlen, gocognit,paralleltest
package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
)

// nolint funlen.
func Test_handler_clientArrives(t *testing.T) {
	t.Parallel()
	clientName := "clientCheckIn"
	type args struct {
		method         string
		uri            string
		requestPayload *RestModel
	}
	tests := []struct {
		name             string
		clientService    Service
		args             args
		expectStatusCode int
		responsePayload  *RestModel
	}{
		{
			name: "fail invalid http.method",
			args: args{
				method: http.MethodPatch,
				uri:    URI,
			},
			expectStatusCode: http.StatusMethodNotAllowed,
		}, {
			name: "fail invalid uri, missing client name",
			args: args{
				method: http.MethodPut,
				uri:    URI,
			},
			expectStatusCode: http.StatusNotFound,
		}, {
			name: "fail invalid body payload, nil body",
			args: args{
				method: http.MethodPut,
				uri:    URI + "/" + clientName,
			},
			expectStatusCode: http.StatusUnprocessableEntity,
		}, {
			name:          "fail checkIn service error",
			clientService: ServiceMock{CheckInErr: internal.ErrStoreCommand},
			args: args{
				method:         http.MethodPut,
				uri:            URI + "/" + clientName,
				requestPayload: &RestModel{GroupSize: 5},
			},
			expectStatusCode: http.StatusInternalServerError,
		}, {
			name:          "Success checkIn Client",
			clientService: ServiceMock{},
			args: args{
				method:         http.MethodPut,
				uri:            URI + "/" + clientName,
				requestPayload: &RestModel{GroupSize: 5},
			},
			expectStatusCode: http.StatusOK,
			responsePayload:  &RestModel{Name: clientName},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			var req *http.Request
			if tt.args.requestPayload != nil {
				var buf bytes.Buffer
				_ = json.NewEncoder(&buf).Encode(tt.args.requestPayload)
				req = httptest.NewRequest(tt.args.method, tt.args.uri, &buf)
			} else {
				req = httptest.NewRequest(tt.args.method, tt.args.uri, nil)
			}

			// invoke
			httpHandler := NewHttpHandler(tt.clientService)
			httpHandler.clientArrives(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.expectStatusCode {
				t.Errorf("GET status = %v, wantStatus %v", res.Status, http.StatusOK)
			}

			if tt.responsePayload != nil {
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("expected error to be nil got %v", err)
				}
				defer res.Body.Close()

				var payload RestModel
				json.Unmarshal(data, &payload)
				if reflect.DeepEqual(payload, tt.responsePayload) {
					t.Errorf("expected body not to be nil: %v", string(data))
				}
			}
		})
	}
}

func Test_handler_clientLeaves(t *testing.T) {
	t.Parallel()
	clientName := "clientCheckOut"
	type args struct {
		method string
		uri    string
	}
	tests := []struct {
		name             string
		clientService    Service
		args             args
		expectStatusCode int
	}{
		{
			name: "fail invalid http.method",
			args: args{
				method: http.MethodPatch,
				uri:    URI,
			},
			expectStatusCode: http.StatusMethodNotAllowed,
		}, {
			name: "fail invalid uri, missing client name",
			args: args{
				method: http.MethodDelete,
				uri:    URI,
			},
			expectStatusCode: http.StatusNotFound,
		}, {
			name:          "fail checkOut service error",
			clientService: ServiceMock{CheckOutErr: internal.ErrStoreCommand},
			args: args{
				method: http.MethodDelete,
				uri:    URI + "/" + clientName,
			},
			expectStatusCode: http.StatusInternalServerError,
		}, {
			name:          "Success checkIn Client",
			clientService: ServiceMock{},
			args: args{
				method: http.MethodDelete,
				uri:    URI + "/" + clientName,
			},
			expectStatusCode: http.StatusNoContent,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tt.args.method, tt.args.uri, nil)

			// invoke
			httpHandler := NewHttpHandler(tt.clientService)
			httpHandler.clientLeaves(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.expectStatusCode {
				t.Errorf("GET status = %v, wantStatus %v", res.Status, http.StatusOK)
			}
		})
	}
}

func Test_handler_clientsInParty(t *testing.T) {
	t.Parallel()
	type args struct {
		method string
		uri    string
	}
	tests := []struct {
		name             string
		clientService    Service
		args             args
		expectStatusCode int
		responsePayload  *[]RestModel
	}{
		{
			name: "fail invalid http.method",
			args: args{
				method: http.MethodPatch,
				uri:    URI,
			},
			expectStatusCode: http.StatusMethodNotAllowed,
		}, {
			name:          "fail FilterByCheckIn service error",
			clientService: ServiceMock{FilterByCheckInErr: internal.ErrStoreCommand},
			args: args{
				method: http.MethodGet,
				uri:    URI,
			},
			expectStatusCode: http.StatusInternalServerError,
		}, {
			name: "Success FilterByCheckIn",
			clientService: ServiceMock{
				FilterByCheckInList: []Client{{
					ID:      1,
					Name:    "name",
					Size:    1,
					CheckIn: "in-time",
					TableID: 1,
				}},
			},
			args: args{
				method: http.MethodGet,
				uri:    URI,
			},
			expectStatusCode: http.StatusOK,
			responsePayload: &[]RestModel{{
				Name:        "name",
				CheckInTime: "in-time",
				GroupSize:   1,
			}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tt.args.method, tt.args.uri, nil)

			// invoke
			httpHandler := NewHttpHandler(tt.clientService)
			httpHandler.clientsInParty(w, req)

			// validate result
			res := w.Result()
			if res.StatusCode != tt.expectStatusCode {
				t.Errorf("GET status = %v, wantStatus %v", res.Status, http.StatusOK)
			}

			if tt.responsePayload != nil {
				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Errorf("expected error to be nil got %v", err)
				}
				defer res.Body.Close()

				var payload RestModel
				json.Unmarshal(data, &payload)
				if reflect.DeepEqual(payload, tt.responsePayload) {
					t.Errorf("expected body not to be nil: %v", string(data))
				}
			}
		})
	}
}
