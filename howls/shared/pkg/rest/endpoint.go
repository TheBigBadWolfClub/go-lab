package rest

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Endpoint interface {
	SubRoutes(r chi.Router)
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
