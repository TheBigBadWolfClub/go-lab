package carddeck

// Rules deal cards constraints
// NumDecks number of decks to deal in (normally the number of players)
// NumCards number of cards per deck, if 0, cards will be split even by all decks
type Rules struct {
	NumDecks int
	NumCards int
}

type Round struct {
	Dead  Deck
	Table Deck
	Decks []Deck
}

type Dealer interface {
	ByPlayer() Round
}

type deal struct {
	rules Rules
	deck  Deck
}

func NewDeal(rules Rules) *deal {
	return &deal{deck: Deck{}, rules: rules}
}

func (d deal) ByPlayer() Round {
	d.deck.Reset()
	d.deck.Shuffle()
	return d.dealBy()
}

// dealBy
// deal nCards per nPlayers
func (d *deal) dealBy() Round {
	subDeckSize := d.rules.decksSize(len(d.deck))
	deals := Round{
		Decks: make([]Deck, d.rules.NumDecks),
	}

	for i := 0; i < d.rules.NumDecks; i++ {
		deals.Decks[i] = make(Deck, subDeckSize)
		start := i * subDeckSize
		subDeck := d.deck[start : start+subDeckSize]
		copy(deals.Decks[i], subDeck)
	}

	lefties := d.deck[d.rules.NumDecks*subDeckSize:]
	deals.Dead = make(Deck, len(lefties))
	copy(deals.Dead, lefties)
	return deals
}

func (r Rules) decksSize(deckSize int) int {
	if r.NumCards == 0 {
		return deckSize / r.NumDecks
	}

	return r.NumCards
}
