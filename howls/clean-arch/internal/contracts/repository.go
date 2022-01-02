// Package contracts
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package contracts

import (
	"github.com/jmoiron/sqlx"
)

const (
	stmList = "SELECT * FROM contracts"
	stmGet  = "SELECT * FROM contracts WHERE type=?"
)

type contractDao struct {
	id        int
	Type      ContractType
	Promotion int
	Max       int
}

type repo struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repo {
	return &repo{db: db}
}

func (r repo) Read(id ContractType) (*Contract, error) {
	dao := contractDao{}
	err := r.db.Get(&dao, stmGet, id)
	if err != nil {
		return nil, err
	}

	return fromDao(dao), nil
}

func (r repo) List() ([]*Contract, error) {
	var daos []contractDao
	err := r.db.Select(&daos, stmList)
	if err != nil {
		return nil, err
	}

	entities := make([]*Contract, 0)
	for _, dao := range daos {
		entities = append(entities, fromDao(dao))
	}
	return entities, nil
}

func fromDao(dao contractDao) *Contract {
	return NewContract(dao.id, dao.Type, dao.Promotion, dao.Max)
}
