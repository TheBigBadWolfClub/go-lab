// Package PowerTools
// Layer: Adapter
// Contains the logic and data representation
// used to interact with frameworks
//
// It's also responsible to convert enterprise entities
// to the one required by frameworks
package PowerTools

import (
	"github.com/jmoiron/sqlx"
)

const (
	stmList   = "SELECT * FROM power_tools"
	stmGet    = "SELECT * FROM power_tools WHERE code=?"
	stmInsert = "INSERT INTO power_tools (code, type,rate) VALUES (:code,:type, :rate)"
	stmUpdate = "UPDATE power_tools SET rate=:rate WHERE code=:code"
	stmDelete = "DELETE FROM power_tools WHERE code=:code"
)

type powerToolsDao struct {
	Code string
	Rate int
	Type string
}

type repo struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repo {
	return &repo{db: db}
}

func (r repo) Read(code Code) (*PowerTool, error) {
	dao := powerToolsDao{}
	err := r.db.Get(&dao, stmGet, code)
	if err != nil {
		return nil, err
	}

	return fromDao(dao), nil
}

func (r repo) List() ([]*PowerTool, error) {
	var daos []powerToolsDao
	err := r.db.Select(&daos, stmList)
	if err != nil {
		return nil, err
	}

	entities := make([]*PowerTool, 0)
	for _, dao := range daos {
		entities = append(entities, fromDao(dao))
	}
	return entities, nil
}

func (r repo) Create(ent *PowerTool) error {
	_, err := r.db.NamedExec(stmInsert,
		map[string]interface{}{
			"code": ent.Code,
			"type": ent.Type,
			"rate": ent.Rate,
		})

	return err
}

func (r repo) Update(ent *PowerTool) error {
	_, err := r.db.NamedExec(stmUpdate,
		map[string]interface{}{
			"rate": ent.Rate,
		})

	return err
}

func (r repo) Delete(code Code) error {
	_, err := r.db.NamedExec(stmDelete,
		map[string]interface{}{
			"code": code,
		})
	return err
}
