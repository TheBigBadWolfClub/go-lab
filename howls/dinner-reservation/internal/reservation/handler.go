// Package reservation supports business logic and domain use cases for entity type Reservation
package reservation

import (
	"encoding/json"
	"errors"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/client"
	"log"
	"net/http"
	"strings"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
)

const URI = "/reservation"

// RestModel the model used to send or receive requests on URI endpoints.
type RestModel struct {
	Name      string `json:"name"`
	Table     int64  `json:"table"`
	GroupSize int    `json:"group_size"`
}

// HttpHandler use cases supported by instance of handler.
type HttpHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

// NewHttpHandler create an instance of handler that supports HttpHandler use cases.
func NewHttpHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// Handler http handler for sub uris of URI and methods of URI.
func (h *handler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getList(w, r)
	case "POST":
		h.addClient(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// AddClient Add a client to the party list.
func (h *handler) addClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	split := strings.Split(r.URL.Path, "/")
	if len(split) != 3 {
		log.Println("Url Param client 'name' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var payload RestModel
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	payload.Name = split[2]
	_, err = h.service.ReserveTable(r.Context(), client.Client{
		TableID: payload.Table,
		Name:    payload.Name,
		Size:    payload.GroupSize,
	})
	if errors.Is(err, internal.ErrNoAvailableSeat) {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(struct{ name string }{payload.Name})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetList Get the reservations list.
func (h *handler) getList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	reservations, err := h.service.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var payload []RestModel
	for _, it := range reservations {
		payload = append(payload, RestModel{
			Name:      it.Name,
			GroupSize: it.Size,
			Table:     it.TableID,
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
