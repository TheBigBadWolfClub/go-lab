package internal

const (
	ErrStoreCommand    DomainError = "fail to execute store command"
	ErrNoAvailableSeat DomainError = "not enough available seats in table"
)

// DomainError provides details about the error that occurred in the domain layer.
type DomainError string

func (d DomainError) Error() string {
	return string(d)
}
