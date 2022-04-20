package table

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/middlewares"
	"github.com/jmoiron/sqlx"
)

const (
	stmInsertf   = `INSERT INTO tables (seats) VALUES %s`
	stmDeleteAll = `DELETE FROM tables`
	stmSelectAll = `SELECT t.id, t.seats, SUM(g.groupSize) as reserved
							FROM tables t
							LEFT OUTER JOIN clients g ON g.tableId = t.id
							GROUP BY t.id`
)

const stmTableAvailableSeats = `SELECT t.seats, SUM(g.groupSize) as reserved
									FROM tables t
									LEFT OUTER JOIN  clients g ON g.tableId = t.id
									WHERE t.id = ? GROUP BY t.id`

const stmTotalAndAvailable = `SELECT SUM(t.seats) as total, SUM(g.groupSize) as reserved
								FROM tables t
								LEFT OUTER JOIN clients g ON g.tableId = t.id`

type Repository interface {
	Save(context.Context, []int) error
	Fetch(context.Context) ([]Table, error)
	Delete(ctx context.Context) error
	TableAvailableSeats(context.Context, int64) (int, error)
	TotalAndAvailableSeats(context.Context) (int, int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, seats []int) error {
	var values []string
	for _, seat := range seats {
		values = append(values, "("+strconv.FormatInt(int64(seat), 10)+")")
	}
	sqlValues := strings.Join(values, ",")
	_, err := r.db.ExecContext(ctx, fmt.Sprintf(stmInsertf, sqlValues))
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "create tables", err)
		return internal.ErrStoreCommand
	}

	return nil
}

func (r *repository) Fetch(ctx context.Context) ([]Table, error) {
	var res []struct {
		Id       int64         `db:"id"`
		Seats    int           `db:"seats"`
		Reserved sql.NullInt64 `db:"reserved"`
	}

	err := r.db.SelectContext(ctx, &res, stmSelectAll)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "fetch list of tables", err)
		return nil, internal.ErrStoreCommand
	}

	list := make([]Table, 0, len(res))
	for _, it := range res {
		list = append(list, Table{
			ID:       it.Id,
			Seats:    it.Seats,
			Reserved: int(it.Reserved.Int64),
		})
	}
	return list, nil
}

func (r *repository) Delete(ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, stmDeleteAll)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "delete tables", err)
		return internal.ErrStoreCommand
	}
	return nil
}

func (r *repository) TableAvailableSeats(ctx context.Context, tableID int64) (int, error) {
	var res []struct {
		Seats    int           `db:"seats"`
		Reserved sql.NullInt64 `db:"reserved"`
	}

	err := r.db.SelectContext(ctx, &res, stmTableAvailableSeats, tableID)
	if err != nil || len(res) != 1 {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "fetch available seats", err)
		return 0, internal.ErrStoreCommand
	}

	return res[0].Seats - int(res[0].Reserved.Int64), nil
}

func (r *repository) TotalAndAvailableSeats(ctx context.Context) (int, int, error) {
	var res struct {
		Total    int `db:"total"`
		Reserved int `db:"reserved"`
	}

	err := r.db.SelectContext(ctx, &res, stmTotalAndAvailable)
	if err != nil {
		log.Printf("%s::%s::%s", ctx.Value(middlewares.RIDKey), "fetch total and available seats", err)
		return 0, 0, internal.ErrStoreCommand
	}

	return res.Total, res.Reserved, nil
}
