package routes

import (
	"github.com/davisenra/papestash/api/handlers"
	"github.com/davisenra/papestash/api/repository"
	"github.com/davisenra/papestash/internal/context"
)

func RegisterRoutes(appCtx *context.AppContext) {
	wallHandler := handlers.WallpaperHandler{
		WallRepo: repository.NewWallpaperRepository(appCtx.Database.Connection),
	}

	api := appCtx.HttpServer.Group("/api/v1")
	api.POST("/login", handlers.Login)
	api.GET("/wallpapers", wallHandler.GetWallpapers)
	api.POST("/wallpapers", wallHandler.UploadWallpaper)
}
