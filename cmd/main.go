package main

import (
	"fmt"
	"log"

	"github.com/davisenra/papestash/api/routes"
	"github.com/davisenra/papestash/internal/context"
	"github.com/davisenra/papestash/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	appConfig, err := context.LoadAppConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabase(appConfig.DatabasePath)

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.MaxMultipartMemory = 16 << 20 // 16 MiB
	r.SetTrustedProxies(nil)        // see: https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies
	gin.DisableConsoleColor()

	appContext := context.AppContext{
		HttpServer: r,
		Database:   db,
	}

	routes.RegisterRoutes(&appContext)

	appHost := fmt.Sprintf(":%d", appConfig.AppPort)

	if err := r.Run(appHost); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
