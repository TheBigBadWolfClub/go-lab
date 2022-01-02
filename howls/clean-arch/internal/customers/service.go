// Package customers
// Layer: Application Business
// Use cases orchestrate the flow of data to and from the entities,
// and direct those entities to use their enterprise wide business rules
// to achieve the goals of the use case.
package customers

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/contracts"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) Get(id internal.ID) (*Customer, error) {
	return s.repo.Read(id)
}

func (s service) List() ([]*Customer, error) {
	return s.repo.List()
}

func (s service) Add(name string, contract string) (*Customer, error) {
	contractType := contracts.ContractType(contract)
	customer := NewCustomer(name, contractType)
	if err := customer.Validate(); err != nil {
		return nil, err
	}

	id, err := s.repo.Create(customer)
	if err != nil {
		return nil, err
	}

	return s.repo.Read(id)
}

func (s service) Update(ent *Customer) error {
	if err := ent.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ent)
}

func (s service) Delete(id internal.ID) error {
	return s.repo.Delete(id)
}
