package handlers

import (
	"net/http"

	"github.com/davisenra/papestash/api/repository"

	"github.com/gin-gonic/gin"
)

type WallpaperHandler struct {
	WallRepo *repository.WallpaperRepository
}

func (h *WallpaperHandler) GetWallpapers(c *gin.Context) {
	wallpapers, err := h.WallRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wallpapers,
	})
}

func (h *WallpaperHandler) UploadWallpaper(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TO BE IMPLEMENTED",
	})
}
