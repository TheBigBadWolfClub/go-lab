package client

import (
	"context"
	"log"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/middlewares"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
)

// Client the entity used in domain and business use cases.
type Client struct {
	ID      int64
	Name    string
	Size    int
	CheckIn string
	TableID int64
}

// Service use cases supported by client entity.
type Service interface {
	Create(context.Context, Client) (int64, error)
	List(context.Context) ([]Client, error)
	CheckIn(context.Context, string, int) error
	CheckOut(context.Context, string) error
	FilterByCheckIn(context.Context) ([]Client, error)
}

type service struct {
	store        Repository
	tableService table.Service
}

// NewService instance of client Service, that orchestrates use-cases related to client entity.
func NewService(store Repository, tableService table.Service) *service {
	return &service{
		store:        store,
		tableService: tableService,
	}
}

// Create save a client into permanent storage.
func (s *service) Create(ctx context.Context, client Client) (int64, error) {
	return s.store.Save(ctx, client)
}

// List Get all client as that will attend to the party.
func (s *service) List(ctx context.Context) ([]Client, error) {
	return s.store.FetchAll(ctx)
}

// CheckIn a client arrived into the party.
func (s *service) CheckIn(ctx context.Context, name string, size int) error {
	findClient, err := s.store.Get(ctx, name)
	if err != nil {
		return err
	}
	findClient.Size = size

	availableSeats, err := s.tableService.AvailableSeats(ctx, findClient.TableID)
	if err != nil {
		return err
	}

	if availableSeats < findClient.Size {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "checkIn client", internal.ErrNoAvailableSeat)
		return internal.ErrNoAvailableSeat
	}

	return s.store.UpdateCheckIn(ctx, findClient)
}

// CheckOut a client leaves the party.
func (s *service) CheckOut(ctx context.Context, name string) error {
	return s.store.Delete(ctx, name)
}

// FilterByCheckIn list all client currently in the party.
func (s *service) FilterByCheckIn(ctx context.Context) ([]Client, error) {
	return s.store.FetchCheckedIn(ctx)
}
