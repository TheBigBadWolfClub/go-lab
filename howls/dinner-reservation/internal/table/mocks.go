package table

import (
	"context"
)

type ServiceMock struct {
	AvailableSeatsError error
	AvailableSeatsGet   int
}

func (m ServiceMock) Add(ctx context.Context, ints []int) error {
	// TODO implement me
	panic("implement me")
}

func (m ServiceMock) Delete(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (m ServiceMock) List(ctx context.Context) ([]Table, error) {
	// TODO implement me
	panic("implement me")
}

func (m ServiceMock) AvailableSeats(ctx context.Context, i int64) (int, error) {
	return m.AvailableSeatsGet, m.AvailableSeatsError
}
