// nolint testpackage, funlen
package reservation

import (
	"context"
	"errors"
	"testing"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/client"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
)

func Test_service_ReserveTable(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		clientService client.Service
		tableService  table.Service
		client        client.Client
		want          int64
		wantErr       error
	}{
		{
			name:          "fail, tableService.AvailableSeats error",
			clientService: client.ServiceMock{},
			tableService:  table.ServiceMock{AvailableSeatsError: internal.ErrStoreCommand},
			client:        client.Client{},
			want:          0,
			wantErr:       internal.ErrStoreCommand,
		}, {
			name:          "fail, table has no seats available",
			clientService: client.ServiceMock{},
			tableService:  table.ServiceMock{AvailableSeatsGet: 10},
			client:        client.Client{Size: 12},
			want:          0,
			wantErr:       internal.ErrNoAvailableSeat,
		}, {
			name:          "fail, clientService.Create  error",
			clientService: client.ServiceMock{CreateErr: internal.ErrStoreCommand},
			tableService:  table.ServiceMock{AvailableSeatsGet: 10},
			client:        client.Client{Size: 3},
			want:          0,
			wantErr:       internal.ErrStoreCommand,
		}, {
			name:          "fail, clientService.Create  error",
			clientService: client.ServiceMock{CreateValue: 2},
			tableService:  table.ServiceMock{AvailableSeatsGet: 10},
			client:        client.Client{Size: 3},
			want:          2,
			wantErr:       nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := NewService(tt.clientService, tt.tableService)
			got, err := s.ReserveTable(context.Background(), tt.client)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("unexpected error, got: %v, exp: %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReserveTable() got = %v, want %v", got, tt.want)
			}
		})
	}
}
