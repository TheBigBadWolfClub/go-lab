// nolint  funlen, gocognit,paralleltest
package client

import (
	"context"
	"testing"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
)

func Test_service_CheckIn(t *testing.T) {
	tests := []struct {
		name         string
		store        Repository
		tableService table.Service
		clientSize   int
		wantErr      bool
	}{
		{
			name: "fail to checkIn, get client error",
			store: repositoryMock{
				GetError: internal.ErrStoreCommand,
			},
			tableService: nil,
			wantErr:      true,
		}, {
			name:  "fail to checkIn, fail to get available seats",
			store: repositoryMock{GetClient: Client{}},
			tableService: table.ServiceMock{
				AvailableSeatsError: internal.ErrStoreCommand,
			},
			wantErr: true,
		}, {
			name:         "fail to checkIn, client list more than available seats",
			store:        repositoryMock{GetClient: Client{}},
			tableService: table.ServiceMock{AvailableSeatsGet: 5},
			clientSize:   10,
			wantErr:      true,
		}, {
			name:         "fail to checkIn, UpdateCheckIn store error",
			store:        repositoryMock{GetClient: Client{}, UpdateCheckInError: internal.ErrStoreCommand},
			tableService: table.ServiceMock{AvailableSeatsGet: 100},
			clientSize:   10,
			wantErr:      true,
		}, {
			name:         "success checkIn",
			store:        repositoryMock{GetClient: Client{}},
			tableService: table.ServiceMock{AvailableSeatsGet: 100},
			clientSize:   10,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := NewService(tt.store, tt.tableService)
			if err := s.CheckIn(context.Background(), "name", tt.clientSize); (err != nil) != tt.wantErr {
				t.Errorf("CheckIn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
