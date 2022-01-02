// Package customers
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package customers

import (
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/contracts"
	"github.com/jmoiron/sqlx"
)

const (
	stmList   = "SELECT * FROM customers LIMIT 10"
	stmGet    = "SELECT * FROM customers WHERE id=?"
	stmInsert = "INSERT INTO customers (Name,Contract) VALUES (:name,:contract)"
	stmUpdate = "UPDATE customers SET name=:name, contract=:contract WHERE id=:id"
	stmDelete = "DELETE FROM customers WHERE id=:id;"
)

type customerDao struct {
	ID       internal.ID
	Name     string
	Contract contracts.ContractType
}

type repo struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repo {
	return &repo{db: db}
}

func (r repo) Read(id internal.ID) (*Customer, error) {
	dao := customerDao{}
	err := r.db.Get(&dao, stmGet, id)
	if err != nil {
		return nil, err
	}

	return fromDao(dao), nil
}

func (r repo) List() ([]*Customer, error) {
	var daos []customerDao
	err := r.db.Select(&daos, stmList)
	if err != nil {
		return nil, err
	}

	entities := make([]*Customer, 0)
	for _, dao := range daos {
		entities = append(entities, fromDao(dao))
	}
	return entities, nil
}

func (r repo) Create(ent *Customer) (internal.ID, error) {
	exec, err := r.db.NamedExec(stmInsert,
		map[string]interface{}{
			"name":     ent.Name,
			"contract": ent.ContractType,
		})

	if err != nil {
		return 0, err
	}

	id, err := exec.LastInsertId()
	return internal.ID(id), err
}

func (r repo) Update(ent *Customer) error {
	_, err := r.db.NamedExec(stmUpdate,
		map[string]interface{}{
			"name":     ent.Name,
			"contract": ent.ContractType,
			"id":       ent.ID,
		})

	return err
}

func (r repo) Delete(id internal.ID) error {
	_, err := r.db.NamedExec(stmDelete,
		map[string]interface{}{
			"id": id,
		})
	return err
}

// Store SQL Adapters

func fromDao(dao customerDao) *Customer {
	entity := NewCustomer(dao.Name, dao.Contract)
	entity.ID = dao.ID
	return entity
}

func toDao(domain Customer) *customerDao {
	return &customerDao{
		ID:       domain.ID,
		Name:     domain.Name,
		Contract: domain.ContractType,
	}
}
