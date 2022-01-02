// Package assignment
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package assignment

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
)

type repo struct {
}

func NewRepository() *repo {
	return &repo{}
}

func (r repo) IsToolAssigned(code PowerTools.Code) error {
	panic("implement me")
}

func (r repo) CustomerTotalAssigned(id internal.ID) (int, error) {
	panic("implement me")
}

func (r repo) Read(id internal.ID, code PowerTools.Code) (*Assignment, error) {
	panic("implement me")
}

func (r repo) Unpaid(id internal.ID) ([]*Assignment, error) {
	panic("implement me")
}

func (r repo) Create(ent *Assignment) error {
	panic("implement me")
}

func (r repo) Update(ent *Assignment) error {
	panic("implement me")
}
