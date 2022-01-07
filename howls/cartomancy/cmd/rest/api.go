package main

import (
	"net/http"

	"github.com/TheBigBadWolfClub/go-lab/howls/cartomancy/internal/carddeck"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", api)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Cartomancy Api"))
	})
	_ = http.ListenAndServe(":8081", r)
}

func api(r chi.Router) {
	var deck carddeck.Deck
	deck.Reset()
	handler := carddeck.NewHandler(deck)
	r.Route("/deck", handler.SubRoutes)
}
