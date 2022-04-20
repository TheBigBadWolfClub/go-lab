package main

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/cartomancy/internal/wargame"
)

func main() {
	game := wargame.NewGame()
	game.Setup()
	game.Play()
}
