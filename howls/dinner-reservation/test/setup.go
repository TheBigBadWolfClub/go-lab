package test

import (
	"context"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
)

const UnlimitedTableSeats = 100000

func SetupFindTableId(tables table.Service) int64 {
	findId := func() int64 {
		tablesList, _ := tables.List(context.Background())
		for _, tl := range tablesList {
			if tl.Seats == UnlimitedTableSeats {
				return tl.ID
			}
		}
		return 0
	}
	// find table ID
	tableID := findId()

	// table does not exist create it, and get it
	if tableID == 0 {
		_ = tables.Add(context.Background(), []int{UnlimitedTableSeats})
		tableID = findId()
	}

	return tableID
}
