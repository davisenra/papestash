package routes

import (
	"github.com/davisenra/papestash/api/handlers"
	"github.com/davisenra/papestash/api/repository"
	"github.com/davisenra/papestash/internal/context"
)

func RegisterRoutes(appCtx *context.AppContext) {
	wallpapersHandler := handlers.WallpaperHandler{
		WallRepo: repository.NewWallpaperRepository(appCtx.Database.Connection),
	}

	api := appCtx.HttpServer.Group("/api/v1")
	api.POST("/login", handlers.Handle(handlers.Login))
	api.GET("/wallpapers", handlers.Handle(wallpapersHandler.GetWallpapers))
	api.POST("/wallpapers", handlers.Handle(wallpapersHandler.UploadWallpaper))
	api.DELETE("/wallpapers/:id", handlers.Handle(wallpapersHandler.DeleteWallpaper))
}
