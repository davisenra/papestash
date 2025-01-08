package handlers

import (
	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/davisenra/papestash/api/repository"
	"github.com/davisenra/papestash/internal/image"
	"github.com/davisenra/papestash/internal/utils"

	"github.com/gin-gonic/gin"
)

const WALLPAPERS_STORAGE_PATH = "storage/wallpapers"
const THUMBNAILS_STORAGE_PATH = "storage/thumbnails"

type WallpaperHandler struct {
	WallRepo *repository.WallpaperRepository
}

func (h *WallpaperHandler) GetWallpapers(c *gin.Context) error {
	aspectRatio := c.Query("aspect_ratio")
	minSize := c.Query("min_size")
	maxSize := c.Query("max_size")
	pageStr := c.Query("page")
	perPageStr := c.Query("per_page")

	var filters []repository.Filter

	if aspectRatio != "" {
		filters = append(filters, repository.FilterByAspectRatio(aspectRatio))
	}

	if minSize != "" && maxSize != "" {
		minSizeInt, err1 := strconv.Atoi(minSize)
		maxSizeInt, err2 := strconv.Atoi(maxSize)
		if err1 == nil && err2 == nil {
			filters = append(filters, repository.FilterBySize(minSizeInt, maxSizeInt))
		}
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		perPage = 25
	}

	result, err := h.WallRepo.GetAll(page, perPage, filters...)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"wallpapers":  result.Wallpapers,
		"page":        result.Page,
		"per_page":    result.PerPage,
		"total_pages": result.TotalPages,
		"total_count": result.TotalCount,
	})

	return nil
}

func (h *WallpaperHandler) UploadWallpaper(c *gin.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	filename := utils.SluggifyFile(file.Filename)

	destinationPath := fmt.Sprintf("%s/%s", WALLPAPERS_STORAGE_PATH, filename)
	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		return fmt.Errorf("failed to save uploaded file: %w", err)
	}

	processor, err := image.NewImageProcessor(destinationPath)
	if err != nil {
		return fmt.Errorf("failed to initialize image processor: %w", err)
	}

	thumbnail, err := processor.GenerateThumbnail(240)
	if err != nil {
		return fmt.Errorf("failed to generate thumbnail: %w", err)
	}

	thumbnailFilename := fmt.Sprintf("thumb_%s", filename)
	thumbnailDestinationPath := fmt.Sprintf("%s/%s", THUMBNAILS_STORAGE_PATH, thumbnailFilename)
	thumbFile, err := os.Create(thumbnailDestinationPath)
	if err != nil {
		return fmt.Errorf("failed to create thumbnail file: %w", err)
	}
	defer thumbFile.Close()

	if err := jpeg.Encode(thumbFile, thumbnail, &jpeg.Options{Quality: 90}); err != nil {
		return fmt.Errorf("failed to encode JPEG thumbnail: %w", err)
	}

	_, err = h.WallRepo.Create(repository.Wallpaper{
		Name:              file.Filename,
		Path:              destinationPath,
		ThumbnailPath:     thumbnailDestinationPath,
		Height:            processor.Height(),
		Width:             processor.Width(),
		AspectRatio:       processor.AspectRatio(),
		SizeInBytes:       processor.Size(),
		MostFrequentColor: processor.MostFrequentColor(image.DEFAULT_DOWNSAMPLE_FACTOR),
		CreatedAt:         time.Now(),
	})
	if err != nil {
		return fmt.Errorf("failed to save wallpaper: %w", err)
	}

	c.Status(http.StatusCreated)

	return nil
}

func (h *WallpaperHandler) DeleteWallpaper(c *gin.Context) error {
	c.Status(http.StatusNoContent)

	wallpaperId := c.Param("id")
	wallpaperIdAsInt, err := strconv.Atoi(wallpaperId)
	if err != nil {
		return err
	}

	wallpaper, err := h.WallRepo.GetById(wallpaperIdAsInt)

	if err != nil {
		return nil
	}

	if err := h.WallRepo.Delete(wallpaperIdAsInt); err != nil {
		return nil
	}

	if err := os.Remove(wallpaper.Path); err != nil {
		return nil
	}

	if err := os.Remove(wallpaper.ThumbnailPath); err != nil {
		return nil
	}

	return nil
}
