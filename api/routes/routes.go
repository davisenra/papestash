package routes

import (
	"github.com/davisenra/papestash/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.POST("/login", handlers.Login)
	api.GET("/wallpapers", handlers.GetWallpapers)
	api.POST("/wallpapers", handlers.UploadWallpaper)
}
