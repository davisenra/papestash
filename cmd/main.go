package main

import (
	"log"

	"github.com/davisenra/papestash/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	r := gin.Default()
	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
