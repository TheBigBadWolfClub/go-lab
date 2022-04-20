package client

import (
	"context"
)

type ServiceMock struct {
	CheckInErr          error
	CheckOutErr         error
	FilterByCheckInErr  error
	FilterByCheckInList []Client
	CreateErr           error
	CreateValue         int64
}

func (s ServiceMock) Create(ctx context.Context, client Client) (int64, error) {
	return s.CreateValue, s.CreateErr
}

func (s ServiceMock) List(ctx context.Context) ([]Client, error) {
	// TODO implement me
	panic("implement me")
}

func (s ServiceMock) CheckIn(ctx context.Context, s2 string, i int) error {
	return s.CheckInErr
}

func (s ServiceMock) CheckOut(ctx context.Context, s2 string) error {
	return s.CheckOutErr
}

func (s ServiceMock) FilterByCheckIn(ctx context.Context) ([]Client, error) {
	return s.FilterByCheckInList, s.FilterByCheckInErr
}

type repositoryMock struct {
	GetError           error
	GetClient          Client
	UpdateCheckInError error
}

func (r repositoryMock) Save(ctx context.Context, client Client) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) FetchAll(ctx context.Context) ([]Client, error) {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) Delete(ctx context.Context, s string) error {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) UpdateCheckIn(ctx context.Context, client Client) error {
	return r.UpdateCheckInError
}

func (r repositoryMock) FetchCheckedIn(ctx context.Context) ([]Client, error) {
	// TODO implement me
	panic("implement me")
}

func (r repositoryMock) Get(ctx context.Context, s string) (Client, error) {
	return r.GetClient, r.GetError
}
