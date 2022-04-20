package carddeck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Deck []Card

// Reset
// set deck to origin, all cards and ordered
func (d *Deck) Reset() {
	d.Empty()
	for _, s := range SuitID.List(SuitID{}) {
		for _, c := range CardID.List(CardID{}) {
			*d = append(*d, NewCard(c, s))
		}
	}
}

func (d *Deck) Empty() {
	*d = make([]Card, 0, 52)
}

func (d Deck) IsEmpty() bool {
	return len(d) == 0
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		n := r.Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}
}

func (d *Deck) Pop() Card {
	c := (*d)[0]
	*d = (*d)[1:]
	return c
}

func (d *Deck) Push(card []Card) {
	*d = append(card, *d...)
}

func (d *Deck) PopBottom() Card {
	c := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return c
}

func (d *Deck) PushBottom(card []Card) {
	*d = append(*d, card...)
}

func (d *Deck) PopIndex(index int) Card {
	c := (*d)[index]
	copy((*d)[index:], (*d)[index+1:]) // Shift a[i+1:] left one index.
	*d = (*d)[:len(*d)-1]              // Truncate slice.
	return c
}

func (d *Deck) PopRandom() Card {
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(51)
	c := (*d)[pos]
	copy((*d)[pos:], (*d)[pos+1:]) // Shift a[i+1:] left one index.
	*d = (*d)[:len(*d)-1]          // Truncate slice.
	return c
}

// Find
// return the index of the card in slice or -1 if not found
func (d Deck) Find(card CardMeta, suit SuitMeta) int {
	for i, c := range d {
		if c.Is(card, suit) {
			return i
		}
	}
	return -1
}

// Has
// true if exists in deck
func (d Deck) Has(card CardMeta, suit SuitMeta) bool {
	return d.Find(card, suit) > -1
}

// BySuit
// deal cards by Suit
func (d *Deck) BySuit() []Deck {
	decks := make([]Deck, Card{}.NumOfSuits())
	for _, c := range *d {
		rank, _ := strconv.Atoi(string(c.suiteID().rank))
		id := rank - 1
		decks[id] = append(decks[id], c)
	}

	d.Empty()
	return decks
}

// QuerySuite
// deal cards by Suit, deck is left unchanged
func (d Deck) QuerySuite(suit SuitStr) (Deck, error) {
	suiteMeta, ok := suit.metadata()
	if !ok {
		return Deck{}, fmt.Errorf("not found")
	}

	var cards Deck
	for _, c := range d {
		if c.SuitMeta == suiteMeta {
			cards = append(cards, c)
		}
	}
	return cards, nil
}

// QueryCard
// query a specific cards, deck is left unchanged
func (d Deck) QueryCard(card CardStr, suit SuitStr) (Card, error) {
	cardMeta, ok := card.metadata()
	if !ok {
		return Card{}, fmt.Errorf("not found")
	}

	suiteMeta, ok := suit.metadata()
	if !ok {
		return Card{}, fmt.Errorf("not found")
	}

	findIndex := d.Find(cardMeta, suiteMeta)
	if findIndex < 0 {
		return Card{}, fmt.Errorf("not found")
	}

	return d[findIndex], nil
}
