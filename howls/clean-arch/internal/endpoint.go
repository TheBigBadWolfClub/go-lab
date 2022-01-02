package internal

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Endpoint interface {
	SubRoutes(r chi.Router)
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
