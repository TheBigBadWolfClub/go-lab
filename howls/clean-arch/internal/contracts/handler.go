// Package contracts
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package contracts

import (
	"encoding/json"
	"net/http"

	"github.com/TheBigBadWolfClub/go-lab/spells/foundation/pkg/rest"

	"github.com/go-chi/chi/v5"
)

type ContractDto struct {
	ID         string `json:"id"`
	MaxDevices int    `json:"max_devices"`
}

type handler struct {
	service Service
}

func NewHandler(srv Service) rest.Endpoint {
	return &handler{service: srv}
}

func (h *handler) SubRoutes(r chi.Router) {
	r.Get("/", h.list)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
	})
}

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.List()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	dtos := make([]*ContractDto, 0)
	for _, v := range list {
		dtos = append(dtos, toPresentation(v))
	}

	marshal, err := json.Marshal(dtos)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	contractType := ContractType(id)
	entity, err := h.service.Get(contractType)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	marshal, err := json.Marshal(toPresentation(entity))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}

// Presentation Rest Adapters.
func toPresentation(entity *Contract) *ContractDto {
	return &ContractDto{
		ID:         string(entity.Type),
		MaxDevices: entity.MaxDevices,
	}
}
