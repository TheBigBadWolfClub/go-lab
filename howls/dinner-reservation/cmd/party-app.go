package main

import (
	"fmt"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/reservation"
	"log"
	"net/http"

	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/client"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/middlewares"
	"github.com/TheBigBadWolfClub/go-lab/howls/dinner-reservation/internal/table"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

const (
	drive    = "mysql"
	addr     = "localhost"
	username = "root"
	password = "toor"
	database = "party"
	port     = 3306

	listenAddr = ":8080"
)

func main() {
	// init mysql.
	db := Open()
	defer db.Close()

	// Table
	tableRepo := table.NewRepository(db)
	tableService := table.NewService(tableRepo)
	tableHandler := table.NewHttpHandler(tableService)
	http.HandleFunc(table.URI, middlewares.RequestIDHandler(tableHandler.Handler))

	// client
	clientRepo := client.NewRepository(db)
	clientService := client.NewService(clientRepo, tableService)
	clientHandler := client.NewHttpHandler(clientService)
	http.HandleFunc(client.URI, middlewares.RequestIDHandler(clientHandler.Handler))

	// reservations
	reservationService := reservation.NewService(clientService, tableService)
	reservationHandler := reservation.NewHttpHandler(reservationService)
	http.HandleFunc(reservation.URI, middlewares.RequestIDHandler(reservationHandler.Handler))

	// ping
	http.ListenAndServe(listenAddr, nil)
}

// Open creates a new connection pool to SQL database
// for the given configuration.
func Open() *sqlx.DB {

	db, err := sqlx.Open(drive, fmt.Sprintf("%s:%s@/%s", username, password, database))
	if err != nil {
		log.Fatalf("fail to connect to db: %v", err)
	}

	return db
}
