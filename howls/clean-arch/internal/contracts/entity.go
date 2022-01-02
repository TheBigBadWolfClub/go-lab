// Package contracts
// Layer: Enterprise Business
// Entities encapsulate Enterprise business rules.
// They are the least likely to change
// when something external changes.
package contracts

import "errors"

type ContractType string

const (
	Premium  ContractType = "premium"
	Standard ContractType = "standard"
	Basic    ContractType = "basic"
	Trial    ContractType = "trial"
)

func (lt ContractType) IsValid() error {
	switch lt {
	case Premium, Standard, Basic, Trial:
		return nil
	}
	return errors.New("invalid contract type")
}

type Contract struct {
	id         int
	Type       ContractType
	Discount   int
	MaxDevices int
}

func NewContract(id int, cType ContractType, discount, maxDevices int) *Contract {
	return &Contract{
		id:         id,
		Type:       cType,
		Discount:   discount,
		MaxDevices: maxDevices,
	}
}
