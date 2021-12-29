package deck

import (
	"encoding/json"
	"github.com/TheBigBadWolfClub/go-lab/spells/foundation/pkg/rest"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type handler struct {
	deck Deck
}

func NewHandler(deck Deck) rest.Endpoint {
	return &handler{deck: deck}
}

func (h *handler) SubRoutes(r chi.Router) {
	r.Get("/", h.full)
	r.Get("/shuffle", h.shuffle)
	r.Get("/deal", h.deal)

	r.Route("/suit/{id}", func(r chi.Router) {
		r.Use(h.validateSuit)
		r.Get("/", h.suite)
	})

	r.Route("/card", func(r chi.Router) {
		r.Get("/search", h.queryCard)
	})
}

func (h *handler) full(w http.ResponseWriter, _ *http.Request) {

	marshal, err := json.Marshal(h.deck)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(w, http.StatusText(er), er)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}

func (h *handler) suite(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	suite, err := h.deck.bySuit(id)
	if err != nil {
		er := http.StatusNotFound
		http.Error(w, http.StatusText(er), er)
		return
	}

	marshal, err := json.Marshal(suite)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(w, http.StatusText(er), er)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}

func (h *handler) validateSuit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if !SuitStr(id).Valid() {
			er := http.StatusNotFound
			http.Error(w, http.StatusText(er), er)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *handler) queryCard(w http.ResponseWriter, r *http.Request) {
	card := r.URL.Query().Get("card")
	suit := r.URL.Query().Get("suit")

	if !SuitStr(suit).Valid() || !CardStr(card).Valid() {
		er := http.StatusNotFound
		http.Error(w, http.StatusText(er), er)
		return
	}

	queryCard, err := h.deck.queryCard(CardStr(card), SuitStr(suit))
	if err != nil {
		er := http.StatusNotFound
		http.Error(w, http.StatusText(er), er)
		return
	}

	marshal, err := json.Marshal(queryCard)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(w, http.StatusText(er), er)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}

func (h *handler) shuffle(w http.ResponseWriter, _ *http.Request) {
	shuffle := h.deck.shuffle()
	marshal, err := json.Marshal(shuffle)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(w, http.StatusText(er), er)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}

func (h *handler) deal(w http.ResponseWriter, r *http.Request) {

	nPlayersStr := r.URL.Query().Get("players")
	nPlayers, err := strconv.Atoi(nPlayersStr)
	if err != nil {
		er := http.StatusConflict
		http.Error(w, "invalid query param: players (n of players)", er)
		return
	}

	nCardsStr := r.URL.Query().Get("cards")
	if nCardsStr == "" {
		nCardsStr = string(rune(len(h.deck.Cards) / nPlayers))
	}

	nCards, err := strconv.Atoi(nPlayersStr)
	if err != nil {
		er := http.StatusConflict
		http.Error(w, "invalid query param: cards (n of cards per player)", er)
		return
	}

	deals := h.deck.deal(nPlayers, nCards)
	marshal, err := json.Marshal(deals)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(w, http.StatusText(er), er)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(marshal)
}
