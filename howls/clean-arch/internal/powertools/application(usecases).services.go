// Package PowerTools
// Layer: Application Business
// Use cases orchestrate the flow of data to and from the entities,
// and direct those entities to use their enterprise wide business rules
// to achieve the goals of the use case.
package PowerTools

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) Get(id Code) (*PowerTool, error) {
	return s.repo.Read(id)
}

func (s service) List() ([]*PowerTool, error) {
	return s.repo.List()
}

func (s service) Add(id Code, toolType string, rate int) (*PowerTool, error) {
	entity := NewPowerTool(id, toolType, rate)

	if err := entity.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (s service) Update(ent *PowerTool) error {
	if err := ent.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ent)
}

func (s service) Delete(id Code) error {
	return s.repo.Delete(id)
}
