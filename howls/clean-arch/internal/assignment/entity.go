// Package assignment
// Layer: Enterprise Business
// Entities encapsulate Enterprise business rules.
// They are the least likely to change
// when something external changes.
package assignment

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
	"time"
)

type Assignment struct {
	Tool       PowerTools.Code
	Customer   internal.ID
	start      string
	end        string
	liquidated bool
}

func NewAssignment(code PowerTools.Code, customer internal.ID) *Assignment {
	return &Assignment{
		Tool:     code,
		Customer: customer,
	}
}

// time.Now().Format("2006-01-02 15:04:05")
func (a *Assignment) Start(time string) *Assignment {
	a.start = time
	return a
}

func (a *Assignment) StartLease() string {
	return a.start
}

func (a *Assignment) End(time string) *Assignment {
	a.end = time
	return a
}
func (a *Assignment) EndLease() string {
	return a.end
}

func (a *Assignment) TotalTime() float64 {
	end, _ := time.Parse(a.end, "2006-01-02 15:04:05")
	start, _ := time.Parse(a.end, "2006-01-02 15:04:05")
	return end.Sub(start).Hours()
}

func (a *Assignment) Liquidate(value bool) *Assignment {
	a.liquidated = value
	return a
}

func (a *Assignment) IsLiquidated() bool {
	return a.liquidated
}
