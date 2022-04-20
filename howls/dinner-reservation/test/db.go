package test

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

const (
	drive    = "mysql"
	addr     = "localhost"
	username = "root"
	password = "toor"
	database = "party"
	port     = 3306
)

// DBOpen creates a new connection pool to SQL database
// for the given configuration.
func DBOpen() *sqlx.DB {
	connStr := fmt.Sprintf("%s:%s@/%s", username, password, database)
	db, err := sqlx.Open(drive, connStr)
	if err != nil {
		panic(fmt.Errorf("fail to connect to db: %v", err))
	}
	return db
}
