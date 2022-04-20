// Package client supports business logic and domain use cases for entity type Client
// also contains the ports and adapter to external frameworks (rest and store)
package client

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
)

const URI = "/client"

// RestModel the model used to send or receive requests on URI endpoints.
type RestModel struct {
	Name        string `json:"name"`
	CheckInTime string `json:"check_in_time"`
	GroupSize   int    `json:"group_size"`
}

// HttpHandler use cases supported by instance of handler.
type HttpHandler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	clientService Service
}

// NewHttpHandler create an instance of handler that supports HttpHandler use cases.
func NewHttpHandler(clientService Service) *handler {
	return &handler{
		clientService: clientService,
	}
}

// Handler http handler for sub uris of URI and methods of URI.
func (h *handler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		h.clientArrives(w, r)
	case "DELETE":
		h.clientLeaves(w, r)
	case "GET":
		h.clientsInParty(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *handler) clientArrives(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	split := strings.Split(r.URL.Path, "/")
	if len(split) != 3 {
		log.Println("Url Param client 'name' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var payload RestModel
	err := decoder.Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	name := split[2]
	err = h.clientService.CheckIn(r.Context(), name, payload.GroupSize)
	if errors.Is(err, internal.ErrNoAvailableSeat) {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(
		struct {
			Name string `json:"name"`
		}{name},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *handler) clientLeaves(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	split := strings.Split(r.URL.Path, "/")
	if len(split) != 3 {
		log.Println("Url Param client 'name' is missing")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	name := split[2]
	err := h.clientService.CheckOut(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) clientsInParty(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	clientList, err := h.clientService.FilterByCheckIn(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var payload []RestModel
	for _, it := range clientList {
		payload = append(payload, RestModel{
			Name:        it.Name,
			GroupSize:   it.Size,
			CheckInTime: it.CheckIn,
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
