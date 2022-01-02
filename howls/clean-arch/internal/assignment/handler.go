// Package assignment
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package assignment

import (
	"encoding/json"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type AssignDto struct {
	Code string `json:"tool_code" validate:"required"`
	Id   int64  `json:"customer_id" validate:"required"`
}

func NewHandler(service Service) internal.Endpoint {
	return &handler{srv: service}
}

type handler struct {
	srv Service
}

func (h handler) SubRoutes(r chi.Router) {
	r.Get("/", h.assign)
}

func (h handler) assign(w http.ResponseWriter, r *http.Request) {
	var dto AssignDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	err = h.srv.Assign(internal.ID(dto.Id), PowerTools.Code(dto.Code))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h handler) unAssign(w http.ResponseWriter, r *http.Request) {
	var dto AssignDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	err = h.srv.UnAssign(internal.ID(dto.Id), PowerTools.Code(dto.Code))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
