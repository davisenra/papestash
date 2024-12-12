package handlers

import (
	"fmt"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/davisenra/papestash/api/repository"
	"github.com/davisenra/papestash/internal/image"

	"github.com/gin-gonic/gin"
)

const WALLPAPERS_STORAGE_PATH = "./storage/wallpapers"
const THUMBNAILS_STORAGE_PATH = "./storage/thumbnails"

type WallpaperHandler struct {
	WallRepo *repository.WallpaperRepository
}

func (h *WallpaperHandler) GetWallpapers(c *gin.Context) error {
	wallpapers, err := h.WallRepo.GetAll()

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"wallpapers": wallpapers,
	})

	return nil
}

func (h *WallpaperHandler) UploadWallpaper(c *gin.Context) error {
	file, err := c.FormFile("file")

	if err != nil {
		return err
	}

	destinationPath := fmt.Sprintf("%s/%s", WALLPAPERS_STORAGE_PATH, file.Filename)

	c.SaveUploadedFile(file, destinationPath)
	processor, _ := image.NewImageProcessor(destinationPath)
	thumbnail, err := processor.GenerateThumbnail(240)

	if err != nil {
		return err
	}

	thumbnailDestinationPath := fmt.Sprintf("%s/%s", THUMBNAILS_STORAGE_PATH, file.Filename)
	thumbFile, _ := os.Create(thumbnailDestinationPath)
	jpeg.Encode(thumbFile, thumbnail, &jpeg.Options{Quality: 90})

	c.JSON(http.StatusOK, gin.H{
		"aspectRatio":       processor.GetAspectRatio(),
		"height":            processor.GetWidth(),
		"width":             processor.GetWidth(),
		"extension":         processor.GetExtension(),
		"size":              processor.GetSize(),
		"mostFrequentColor": processor.GetMostFrequentColor(8),
	})

	return nil
}
