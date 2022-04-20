// Package table supports business logic and domain use cases for entity type Table
// also contains the ports and adapter to external frameworks (rest and store)
package table

import (
	"encoding/json"
	"net/http"
)

const URI = "/tables"

// RestModel the model used to send or receive requests on URI endpoints.
type RestModel struct {
	ID        int64 `json:"id"`
	Seats     int   `json:"seats"`
	Available int   `json:"available"`
}

// HttpHandler use cases supported by instance of handler.
type HttpHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

// NewHttpHandler create an instance of handler that supports HttpHandler use cases.
func NewHttpHandler(srv Service) *handler {
	return &handler{
		service: srv,
	}
}

// Handler http handler for sub uris of URI and methods of URI.
func (h *handler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.addTables(w, r)
	case "DELETE":
		h.deleteTables(w, r)
	case "GET":
		h.getTables(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *handler) addTables(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var payload []int
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = h.service.Add(r.Context(), payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) deleteTables(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := h.service.Delete(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) getTables(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	list, err := h.service.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var payload []RestModel
	for _, it := range list {
		payload = append(payload, RestModel{
			ID:        it.ID,
			Seats:     it.Seats,
			Available: it.Seats - it.Reserved,
		})
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
