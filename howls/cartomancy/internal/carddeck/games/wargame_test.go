package games

import (
	"fmt"
	"github.com/TheBigBadWolfClub/go-lab/howls/cartomancy/internal/carddeck"
	"testing"
)

func TestPlayWarGame(t *testing.T) {
	var deck1 carddeck.Deck
	deck1.Reset()
	deck1.Shuffle()
	players := deck1.ByPlayers(2, 26)

	play(players)

}

func play(players carddeck.Deal) {

	var inTable []carddeck.Card
	rounds := 0
	for {

		rounds++
		if rounds > 100 {
			players.Players[1].Shuffle()
			players.Players[0].Shuffle()
			fmt.Println("shuffle -------------------------------------")
			rounds = 0
		}
		if len(players.Players[1]) == 0 {
			fmt.Printf("P1 Win\n")
			fmt.Println(inTable)
			return
		}

		if len(players.Players[0]) == 0 {
			fmt.Printf("P2 Win\n")
			fmt.Println(inTable)

			return
		}

		p1 := players.Players[0].Pop()
		p2 := players.Players[1].Pop()
		inTable = append(inTable, p1, p2)

		fmt.Printf("p1[%v]p2[%v] ", len(players.Players[0]), len(players.Players[1]))
		fmt.Printf("%v-%v\n", p1.CardMeta, p2.CardMeta)
		if p1.CompareTo(p2) == 1 {
			players.Players[0].PushBottom(inTable)
			inTable = nil
		}

		if p1.CompareTo(p2) == -1 {
			players.Players[1].PushBottom(inTable)
			inTable = nil
		}
	}
}
