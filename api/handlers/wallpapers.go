package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWallpapers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TO BE IMPLEMENTED",
	})
}

func UploadWallpaper(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TO BE IMPLEMENTED",
	})
}
