// Package contracts
// Layer: Application Business
// Use cases orchestrate the flow of data to and from the entities,
// and direct those entities to use their enterprise wide business rules
// to achieve the goals of the use case.
package contracts

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (u service) Get(id ContractType) (*Contract, error) {
	if err := id.IsValid(); err != nil {
		return nil, err
	}

	return u.repo.Read(id)
}

func (u service) List() ([]*Contract, error) {
	return u.repo.List()
}
