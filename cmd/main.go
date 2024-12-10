package main

import (
	"github.com/davisenra/papestash/api/routes"
	"github.com/davisenra/papestash/internal/context"
	"github.com/davisenra/papestash/internal/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := database.NewDatabase("db.database")

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	gin.DisableConsoleColor()

	appContext := context.AppContext{
		HttpServer: r,
		Database:   db,
	}

	routes.RegisterRoutes(&appContext)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
