package client

import (
	"context"
	"database/sql"
	"log"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/middlewares"
	"github.com/jmoiron/sqlx"
)

const (
	stmInsert           = `INSERT INTO clients (name, groupSize, tableId) VALUES (:name,:groupSize, :tableId)`
	stmSelect           = `SELECT id, name, groupSize, checkIn, tableId FROM clients WHERE name=?`
	stmSelectAll        = `SELECT id, name, groupSize, checkIn FROM clients`
	stmSelectAllCheckIn = `SELECT id, name, groupSize, checkIn, tableId FROM clients WHERE checkIn IS NOT NULL`
	stmDelete           = `DELETE FROM clients WHERE name=:name`
	stmUpdateCheckIn    = `UPDATE  clients SET groupSize = :groupSize, checkIn = NOW() WHERE name = :name`
)

type Repository interface {
	Save(context.Context, Client) (int64, error)
	FetchAll(context.Context) ([]Client, error)
	Delete(context.Context, string) error
	UpdateCheckIn(context.Context, Client) error
	FetchCheckedIn(context.Context) ([]Client, error)
	Get(context.Context, string) (Client, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, client Client) (int64, error) {
	args := map[string]interface{}{
		"name":      client.Name,
		"groupSize": client.Size,
		"tableId":   client.TableID,
	}
	result, err := r.db.NamedExecContext(ctx, stmInsert, args)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "add client", err)
		return 0, internal.ErrStoreCommand
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "add client", err)
		return 0, internal.ErrStoreCommand
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, name string) (Client, error) {
	var result struct {
		Id        int64          `db:"id"`
		Name      string         `db:"name"`
		GroupSize int            `db:"groupSize"`
		CheckIn   sql.NullString `db:"checkIn"`
		TableID   int64          `db:"tableId"`
	}
	err := r.db.GetContext(ctx, &result, stmSelect, name)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "get client", err)
		return Client{}, internal.ErrStoreCommand
	}

	return Client{
		ID:      result.Id,
		Name:    result.Name,
		Size:    result.GroupSize,
		CheckIn: result.CheckIn.String,
		TableID: result.TableID,
	}, nil
}

func (r *repository) Delete(ctx context.Context, name string) error {
	args := map[string]interface{}{
		"name": name,
	}
	_, err := r.db.NamedExecContext(ctx, stmDelete, args)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "delete client", err)
		return internal.ErrStoreCommand
	}

	return nil
}

func (r *repository) UpdateCheckIn(ctx context.Context, client Client) error {
	args := map[string]interface{}{
		"name":      client.Name,
		"groupSize": client.Size,
	}
	_, err := r.db.NamedExecContext(ctx, stmUpdateCheckIn, args)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "checkIn client", err)
		return internal.ErrStoreCommand
	}

	return nil
}

func (r *repository) FetchCheckedIn(ctx context.Context) ([]Client, error) {
	var res []struct {
		Id        int64  `db:"id"`
		Name      string `db:"name"`
		GroupSize int    `db:"groupSize"`
		CheckIn   string `db:"checkIn"`
		TableID   int64  `db:"tableId"`
	}

	err := r.db.SelectContext(ctx, &res, stmSelectAllCheckIn)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "fetch list of clients", err)
		return nil, internal.ErrStoreCommand
	}

	list := make([]Client, 0, len(res))
	for _, it := range res {
		client := Client{
			ID:      it.Id,
			Name:    it.Name,
			Size:    it.GroupSize,
			CheckIn: it.CheckIn,
			TableID: it.TableID,
		}
		list = append(list, client)
	}
	return list, nil
}

func (r *repository) FetchAll(ctx context.Context) ([]Client, error) {
	var result []struct {
		ID      int64          `db:"id"`
		Name    string         `db:"name"`
		Size    int            `db:"groupSize"`
		CheckIn sql.NullString `db:"checkIn"`
	}

	err := r.db.SelectContext(ctx, &result, stmSelectAll)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "fetch list of clients", err)
		return nil, internal.ErrStoreCommand
	}

	var list []Client
	for _, it := range result {
		client := Client{
			ID:      it.ID,
			Name:    it.Name,
			Size:    it.Size,
			CheckIn: it.CheckIn.String,
		}
		list = append(list, client)
	}
	return list, nil
}
