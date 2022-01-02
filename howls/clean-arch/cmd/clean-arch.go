package main

import (
	"fmt"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/assignment"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/billing"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/contracts"
	"github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/customers"
	PowerTools "github.com/TheBigBadWolfClub/go-lab/howls/clean-arch/internal/powertools"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

const (
	sqlUser     = "root"
	sqlPass     = "toor"
	sqlHost     = "localhost"
	sqlPort     = 3306
	sqlDatabase = "cleanDB"
)

func main() {
	db := connectDB()
	router := chi.NewRouter()

	// Contracts
	contractsRepo := contracts.NewRepository(db)
	contractsUseCases := contracts.NewService(contractsRepo)
	contractsHandler := contracts.NewHandler(contractsUseCases)
	router.Route("/contracts", contractsHandler.SubRoutes)

	// Customers
	customersRepo := customers.NewRepository(db)
	customersService := customers.NewService(customersRepo)
	customersHandler := customers.NewHandler(customersService)
	router.Route("/customers", customersHandler.SubRoutes)

	// Tools
	powerToolsRepo := PowerTools.NewRepository(db)
	powerToolsService := PowerTools.NewService(powerToolsRepo)
	powerToolsHandler := PowerTools.NewHandler(powerToolsService)
	router.Route("/power-tools", powerToolsHandler.SubRoutes)

	// assign
	assignRepo := assignment.NewRepository()
	assignService := assignment.NewService(assignRepo, customersRepo, contractsRepo)
	assignHandler := assignment.NewHandler(assignService)
	router.Route("/power-tools/{id}/assignment", assignHandler.SubRoutes)

	// billing
	billingService := billing.NewService(assignRepo, powerToolsRepo, customersRepo, contractsRepo)
	billingHandler := billing.NewHandler(billingService)
	router.Route("/customers/{id}", billingHandler.SubRoutes)

	_ = http.ListenAndServe(":8011", router)
}

func connectDB() *sqlx.DB {

	conStr := fmt.Sprintf("%v:%v@(%v:%d)/%v", sqlUser, sqlPass, sqlHost, sqlPort, sqlDatabase)
	db, err := sqlx.Connect("mysql", conStr)
	if err != nil {
		log.Fatalf("error connecting to DB config file: %v", err)
	}

	return db
}
