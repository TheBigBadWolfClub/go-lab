package wargame

import (
	"fmt"

	"github.com/TheBigBadWolfClub/go-lab/howls/cartomancy/internal/carddeck"
)

type playerAction int

const (
	playCard  playerAction = 1
	cardCount playerAction = 2
)

var ids []string

func init() {
	ids = append(ids, []string{"p1", "p2"}...)
}

type round struct {
	stats  []stats
	hands  []hand
	winner int
}

func (r *round) addHand(h hand) {
	r.hands = append(r.hands, h)
	fmt.Printf("%d: %s \n", h.id, h.card)
}

func (r *round) addStats(s stats) {
	r.stats = append(r.stats, s)
}

func (r *round) findHandWinner() {
	winHand := r.hands[0]
	for i := 1; i < len(r.hands); i++ {
		to := winHand.card.CompareTo(r.hands[i].card)
		if to < 0 {
			winHand = r.hands[i]
		}
	}

	fmt.Printf("win: %d: %s \n", winHand.id, winHand.card)
	r.winner = winHand.id
}

func (r *round) cards() []carddeck.Card {
	var cards []carddeck.Card
	for _, h := range r.hands {
		cards = append(cards, h.card)
	}
	return cards
}

type history []round

func (h *history) add(r round) {
	*h = append(*h, r)
}

type Game struct {
	players Players
	dealer  carddeck.Dealer
	table   *table
}

func NewGame() *Game {
	deal := carddeck.NewDeal(carddeck.Rules{NumDecks: len(ids)})
	return &Game{
		dealer:  deal,
		players: []*Player{},
		table:   &table{},
	}
}

func (g *Game) Play() {
	g.start()
}

func (g *Game) Setup() {
	g.setupPlayers()
	g.setupTable()
}

func (g *Game) setupTable() {
	g.table.playerHands = make(chan hand, len(g.players))
	g.table.playerStats = make(chan stats, len(g.players))
	g.table.players = g.players
	g.table.claimWinner = make(chan int)
}

func (g *Game) setupPlayers() {
	for i, pid := range ids {
		player := Player{
			id:         i,
			name:       pid,
			cards:      carddeck.Deck{},
			actionChan: make(chan playerAction, 1),
			winHand:    make(chan []carddeck.Card, 1),
			table:      g.table,
		}
		g.players = append(g.players, &player)
	}
}

func (g *Game) start() {
	decks := g.dealer.ByPlayer()
	g.players.start(decks)
	g.table.start()

	// wait for winner
	<-g.table.claimWinner
}

type Players []*Player

func (p *Players) askPlayerAction(action playerAction) {
	for _, pl := range *p {
		pl.actionChan <- action
	}
}

func (p *Players) start(decks carddeck.Round) {
	for i, pl := range *p {
		pl.cards.Empty()
		pl.cards.Push(decks.Decks[i])
		pl.initPlay()
	}
}

type Player struct {
	id         int
	name       string
	cards      carddeck.Deck
	actionChan chan playerAction
	winHand    chan []carddeck.Card // number of cards
	table      *table
}

func (p *Player) initPlay() {
	go func() {
		for {
			select {
			case cards := <-p.winHand:
				// prev playerHands res
				p.cards.Push(cards)
			case action := <-p.actionChan:
				switch action {
				case playCard:
					p.table.playerHands <- hand{
						id:   p.id,
						card: p.cards.Pop(),
					}
				case cardCount:
					p.table.playerStats <- stats{
						id:        p.id,
						cardsLeft: len(p.cards),
					}
				}
			}
		}
	}()
}

type table struct {
	playerHands chan hand
	playerStats chan stats
	players     Players
	history     history
	claimWinner chan int
}

func (t *table) start() {
	go func() {
		for {
			var curRound round

			// get card
			t.players.askPlayerAction(playCard)
			for range ids {
				curRound.addHand(<-t.playerHands)
			}

			// find weHaveWinner
			curRound.findHandWinner()

			// send cards to winner
			t.players[curRound.winner].winHand <- curRound.cards()

			// get  count
			t.players.askPlayerAction(cardCount)
			for range ids {
				curRound.addStats(<-t.playerStats)
			}

			// save history
			t.history.add(curRound)

			for _, s := range curRound.stats {
				if s.cardsLeft == 52 {
					fmt.Println("winner: ", s.id)
					t.claimWinner <- s.id
					return
				}
			}
		}
	}()
}

type hand struct {
	id   int
	card carddeck.Card
}

type stats struct {
	id        int
	cardsLeft int
}
