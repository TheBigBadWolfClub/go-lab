// Package PowerTools
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package PowerTools

import (
	"encoding/json"
	"github.com/TheBigBadWolfClub/go-lab/spells/foundation/pkg/rest"
	"net/http"

	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	"github.com/go-chi/chi/v5"
)

type PowerToolDto struct {
	Code string
	Type string
	Rate int
}

type handler struct {
	service Service
}

func NewHandler(srv Service) rest.Endpoint {
	return &handler{service: srv}
}

func (h *handler) SubRoutes(r chi.Router) {
	r.Get("/", h.list)
	r.Post("/", h.post)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(h.validateId)
		r.Get("/", h.get)
		r.Put("/", h.put)
		r.Delete("/", h.delete)
	})
}

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.List()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	dtos := make([]*PowerToolDto, 0)
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

func (h *handler) post(w http.ResponseWriter, r *http.Request) {
	var newData PowerToolDto
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	newEntity, err := h.service.Add(Code(newData.Code), newData.Type, newData.Rate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(toPresentation(newEntity))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	entity, err := h.service.Get(Code(id))
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

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.Delete(Code(id))
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) put(w http.ResponseWriter, r *http.Request) {
	var newData PowerToolDto
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	err = h.service.Update(fromPresentation(newData))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) validateId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := internal.ValidId(id); err != nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
		next.ServeHTTP(w, r)
	})
}
