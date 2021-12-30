package db

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const driverName = "mysql"

func ConnectDB(config *mysql.Config) *sqlx.DB {
	db, err := sqlx.Connect(driverName, config.FormatDSN())
	if err != nil {
		log.Fatalf("error connecting to DB config file: %v", err)
	}

	return db
}

func GetDefaultConfig() *mysql.Config {
	return &mysql.Config{
		User:                 "root",
		Passwd:               "toor",
		Net:                  "tcp",
		Addr:                 "localhost:11306",
		DBName:               "creatures",
		Collation:            "utf8_general_ci",
		AllowNativePasswords: true,
	}
}
