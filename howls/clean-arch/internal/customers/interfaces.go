// Package customers
// Layer: Application Business
// Interfaces define abstractions so layers can be referred
// 1) Inner layers define a set of interface
// 2) Outer layers implement interfaces
// 3) Dependency injection patter is used to provide instances to inner layers
// 4) Inner layers act on outer layers thought this interfaces abstractions
package customers

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
)

//Reader interface
type Reader interface {
	Read(id internal.ID) (*Customer, error)
	List() ([]*Customer, error)
}

//Writer interface
type Writer interface {
	Create(ent *Customer) (internal.ID, error)
	Update(ent *Customer) error
	Delete(id internal.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//Service interface
type Service interface {
	Get(id internal.ID) (*Customer, error)
	List() ([]*Customer, error)
	Add(name string, contractType string) (*Customer, error)
	Update(ent *Customer) error
	Delete(id internal.ID) error
}
