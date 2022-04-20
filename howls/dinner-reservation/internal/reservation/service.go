package reservation

import (
	"context"
	"log"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/client"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/middlewares"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
)

// Service use cases supported by reservation.
type Service interface {
	ReserveTable(context.Context, client.Client) (int64, error)
	List(context.Context) ([]client.Client, error)
}

type service struct {
	clientService client.Service
	tableService  table.Service
}

// NewService instance of reservation Service, that orchestrates use-cases related to reservation entity.
func NewService(clientService client.Service, tableService table.Service) *service {
	return &service{
		clientService: clientService,
		tableService:  tableService,
	}
}

// ReserveTable add a client to list and reserve seats at given table id.
func (s *service) ReserveTable(ctx context.Context, client client.Client) (int64, error) {
	availableSeats, err := s.tableService.AvailableSeats(ctx, client.TableID)
	if err != nil {
		return 0, err
	}

	if availableSeats < client.Size {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "reserve table", internal.ErrNoAvailableSeat)
		return 0, internal.ErrNoAvailableSeat
	}

	return s.clientService.Create(ctx, client)
}

// List gets a list of all client that are in party or have a reservation.
func (s *service) List(ctx context.Context) ([]client.Client, error) {
	return s.clientService.List(ctx)
}
