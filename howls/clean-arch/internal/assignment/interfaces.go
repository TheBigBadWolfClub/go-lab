// Package assignment
// Layer: Application Business
// Interfaces define abstractions so layers can be referred
// 1) Inner layers define a set of interface
// 2) Outer layers implement interfaces
// 3) Dependency injection patter is used to provide instances to inner layers
// 4) Inner layers act on outer layers thought this interfaces abstractions
package assignment

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
)

//Reader interface
type Reader interface {
	IsToolAssigned(code PowerTools.Code) error
	CustomerTotalAssigned(id internal.ID) (int, error)
	Read(id internal.ID, code PowerTools.Code) (*Assignment, error)
	Unpaid(id internal.ID) ([]*Assignment, error)
}

//Writer interface
type Writer interface {
	Create(ent *Assignment) error
	Update(ent *Assignment) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//Service interface
type Service interface {
	Assign(customerId internal.ID, code PowerTools.Code) error
	UnAssign(customerId internal.ID, code PowerTools.Code) error
}
