// Package customers
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package customers

import (
	"encoding/json"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/contracts"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type CustomerDto struct {
	ID       internal.ID `json:"id"`
	Name     string      `json:"name"`
	Contract string      `json:"contract"`
}

type handler struct {
	service Service
}

func NewHandler(srv Service) internal.Endpoint {
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

	dtos := make([]*CustomerDto, 0)
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
	var newData CustomerDto
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	newEntity, err := h.service.Add(newData.Name, newData.Contract)
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
	idStr := chi.URLParam(r, "id")
	id := internal.FromString(idStr)
	entity, err := h.service.Get(id)
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
	idStr := chi.URLParam(r, "id")
	id := internal.FromString(idStr)
	err := h.service.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) put(w http.ResponseWriter, r *http.Request) {
	var newData CustomerDto
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

func toPresentation(v *Customer) *CustomerDto {
	return &CustomerDto{
		ID:       v.ID,
		Name:     v.Name,
		Contract: string(v.ContractType),
	}
}

func fromPresentation(v CustomerDto) *Customer {
	contract := contracts.ContractType(v.Contract)
	customer := NewCustomer(v.Name, contract)
	customer.ID = v.ID
	return customer
}
