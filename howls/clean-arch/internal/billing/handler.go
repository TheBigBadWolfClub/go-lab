package billing

import (
	"encoding/json"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewHandler(service Service) internal.Endpoint {
	return &handler{srv: service}
}

type PayDto struct {
	Code string `json:"tool_code" validate:"required"`
	Id   int64  `json:"customer_id" validate:"required"`
}

type handler struct {
	srv Service
}

func (h handler) SubRoutes(r chi.Router) {
	r.Get("/pay", h.pay)
	r.Get("/total", h.total)
}

func (h handler) pay(w http.ResponseWriter, r *http.Request) {
	var dto PayDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	err = h.srv.Pay(internal.ID(dto.Id), PowerTools.Code(dto.Code))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h handler) total(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id := internal.FromString(idStr)
	total, err := h.srv.Total(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	marshal, err := json.Marshal(struct {
		Total float64
	}{Total: total})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}
