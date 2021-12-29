package main

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/cartomancy/internal/deck"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
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

	handler := deck.NewHandler(deck.NewDeck())
	r.Route("/deck", handler.SubRoutes)

}
