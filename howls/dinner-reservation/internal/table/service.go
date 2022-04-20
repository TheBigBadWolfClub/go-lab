package table

import "context"

// Table the entity used in domain and business use cases.
type Table struct {
	ID       int64
	Seats    int
	Reserved int
}

// Service use cases supported by client entity.
type Service interface {
	Add(context.Context, []int) error
	Delete(context.Context) error
	List(context.Context) ([]Table, error)
	AvailableSeats(context.Context, int64) (int, error)
}

type service struct {
	store Repository
}

// NewService instance of client Service, that orchestrates use-cases related to client entity.
func NewService(store Repository) *service {
	return &service{store: store}
}

// Add adds list of tables to the party.
func (s *service) Add(ctx context.Context, seats []int) error {
	return s.store.Save(ctx, seats)
}

// Delete deletes all tables from the party.
func (s *service) Delete(ctx context.Context) error {
	return s.store.Delete(ctx)
}

// List all tables available at the party.
func (s *service) List(ctx context.Context) ([]Table, error) {
	return s.store.Fetch(ctx)
}

// AvailableSeats total number of seats available.
func (s *service) AvailableSeats(ctx context.Context, id int64) (int, error) {
	return s.store.TableAvailableSeats(ctx, id)
}
