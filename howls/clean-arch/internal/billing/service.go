// Package billing
// Layer: Application Business
// Use cases orchestrate the flow of data to and from the entities,
// and direct those entities to use their enterprise wide business rules
// to achieve the goals of the use case.
package billing

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/assignment"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/contracts"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/customers"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
)

// Service interface.
type Service interface {
	UnpaidList(id internal.ID) ([]*assignment.Assignment, error)
	Pay(id internal.ID, code PowerTools.Code) error
	Total(id internal.ID) (float64, error)
}

type service struct {
	assignRepo    assignment.Repository
	customerRepo  customers.Reader
	contractsRepo contracts.Repository
	toolsRepo     PowerTools.Reader
}

func NewService(assignRepo assignment.Repository, toolsRepo PowerTools.Reader, customerRepo customers.Reader, contractsRepo contracts.Repository) Service {
	return &service{
		assignRepo:    assignRepo,
		toolsRepo:     toolsRepo,
		customerRepo:  customerRepo,
		contractsRepo: contractsRepo,
	}
}

func (s service) UnpaidList(id internal.ID) ([]*assignment.Assignment, error) {
	return s.assignRepo.Unpaid(id)
}

func (s service) Pay(id internal.ID, code PowerTools.Code) error {
	assign, err := s.assignRepo.Read(id, code)
	if err != nil {
		return internal.ErrInconsistentState
	}

	assign.Liquidate(true)
	if err = s.assignRepo.Update(assign); err != nil {
		return internal.ErrFailPayment
	}

	return nil
}

func (s service) Total(id internal.ID) (float64, error) {
	list, err := s.UnpaidList(id)
	if err != nil {
		return 0, internal.ErrInconsistentState
	}

	customer, err := s.customerRepo.Read(id)
	if err != nil {
		return 0, internal.ErrNotFound
	}

	contract, err := s.contractsRepo.Read(customer.ContractType)
	if err != nil {
		return 0, internal.ErrNotFound
	}

	total := float64(0)
	for _, v := range list {
		tool, _ := s.toolsRepo.Read(v.Tool)
		total += calculatePrice(v.TotalTime(), tool.Rate, contract.Discount)
	}

	return total, nil
}

func calculatePrice(time float64, rate int, discount int) float64 {
	return time * float64(rate) * float64(discount/100) //nolint
}
