package main

import (
	"http_server/internal/calculationService"
	"http_server/internal/db"
	"http_server/internal/handlers"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	calcRepo := calculationService.NewCalculationRepository(database)
	calcService := calculationService.NewCalculationService(calcRepo)
	calcHandlers := handlers.NewCalculationHandler(calcService)

	server := echo.New() //creating a server instance
	server.Use(middleware.Logger())
	server.Use(middleware.CORS())

	server.GET("/calculations", calcHandlers.GetCalculations)
	server.POST("/calculations", calcHandlers.PostCalculations)
	server.PATCH("/calculations/:id", calcHandlers.PatchCalculation)
	server.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	server.Start("localhost:8080")

}
