// Package customers
// Layer: Enterprise Business
// Entities encapsulate Enterprise business rules.
// They are the least likely to change
// when something external changes.
package customers

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/contracts"
)

// Customer entity used on the core layer and application layer
type Customer struct {
	ID           internal.ID
	Name         string
	ContractType contracts.ContractType
}

func NewCustomer(name string, cType contracts.ContractType) *Customer {
	return &Customer{
		Name:         name,
		ContractType: cType,
	}
}

func (c *Customer) Validate() error {
	if c.Name == "" || c.ContractType.IsValid() != nil {
		return internal.ErrInvalidEntity
	}
	return nil
}
