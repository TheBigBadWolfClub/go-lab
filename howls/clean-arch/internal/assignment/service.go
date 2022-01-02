// Package assignment
//Layer: Application Business
// Use cases orchestrate the flow of data to and from the entities,
// and direct those entities to use their enterprise wide business rules
// to achieve the goals of the use case.
package assignment

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/contracts"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/customers"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
	"time"
)

type service struct {
	customerRepo  customers.Reader
	repo          Repository
	contractsRepo contracts.Repository
}

func NewService(repo Repository, customers customers.Reader, contracts contracts.Repository) Service {
	return &service{
		repo:          repo,
		customerRepo:  customers,
		contractsRepo: contracts,
	}
}

func (s service) Assign(customerId internal.ID, code PowerTools.Code) error {
	customer, err := s.customerRepo.Read(customerId)
	if err != nil {
		return internal.ErrNotFound
	}

	contract, err := s.contractsRepo.Read(customer.ContractType)
	if err != nil {
		return internal.ErrNotFound
	}

	totalAssigned, err := s.repo.CustomerTotalAssigned(customerId)
	if totalAssigned >= contract.MaxDevices {
		return internal.ErrMaxAllowedReached
	}

	if err = s.repo.IsToolAssigned(code); err != nil {
		return internal.ErrToolUnavailable
	}

	assigned := NewAssignment(code, customerId)
	assigned.Start(time.Now().Format("2006-01-02 15:04:05"))
	return s.repo.Create(assigned)
}

func (s service) UnAssign(customerId internal.ID, code PowerTools.Code) error {

	assign, err := s.repo.Read(customerId, code)
	if err != nil {
		return internal.ErrInconsistentState
	}

	assign.End(time.Now().Format("2006-01-02 15:04:05"))
	return s.repo.Update(assign)
}
