package core

import (
	"math/rand"
	"time"
)

type DiceRoller interface {
	Roll() int
}

type dice struct {
}

func (d dice) Roll() int {
	sleepy := time.Duration(rand.Intn(10))
	time.Sleep(sleepy * time.Millisecond)

	return rand.Intn(6) + 1
}

func NewDiceRoller() DiceRoller {
	return &dice{}
}
