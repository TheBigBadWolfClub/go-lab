package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deals struct {
	Undeal Deck
	Deals  []Deck
}

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	return Deck{Cards: buildDeck()}
}

func buildDeck() []Card {
	var deck []Card
	for _, s := range SuitID.List(SuitID{}) {
		for _, c := range CardID.List(CardID{}) {
			deck = append(deck, NewCard(c, s))
		}
	}
	return deck
}

func (d Deck) bySuit(id string) ([]Card, error) {
	if !SuitStr(id).Valid() {
		return nil, fmt.Errorf("")
	}
	var list []Card
	for _, c := range d.Cards {
		list = append(list, c)
	}
	return list, nil
}

func (d Deck) queryCard(card CardStr, suit SuitStr) (Card, error) {
	cardMeta, ok := card.metadata()
	if !ok {
		return Card{}, fmt.Errorf("not found")
	}

	suiteMeta, ok := suit.metadata()
	if !ok {
		return Card{}, fmt.Errorf("not found")
	}

	newCard := NewCard(cardMeta.cardID(), suiteMeta.suiteID())
	for _, c := range d.Cards {
		if c.Equal(newCard) {
			return c, nil
		}
	}

	return Card{}, fmt.Errorf("not found")
}

func (d Deck) shuffle() Deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, _ := range d.Cards {
		n := r.Intn(len(d.Cards) - 1)
		d.Cards[i], d.Cards[n] = d.Cards[n], d.Cards[i]
	}
	return Deck{Cards: d.Cards}
}

func (d Deck) deal(players, cards int) Deals {

	deals := Deals{
		Undeal: d.shuffle(),
		Deals:  make([]Deck, players),
	}
	for i := 0; i < players; i++ {
		deals.Deals[i] = Deck{Cards: deals.Undeal.Cards[:cards]}
		deals.Undeal.Cards = deals.Undeal.Cards[cards:]
	}

	return deals
}
