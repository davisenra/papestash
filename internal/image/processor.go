package image

import (
	"fmt"
	"image"
	"math"
	"os"
	"path/filepath"
	"sync"

	"github.com/disintegration/imaging"
)

type Processor struct {
	img      image.Image
	filePath string
}

func NewImageProcessor(imagePath string) (*Processor, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return &Processor{img: img, filePath: imagePath}, nil
}

func (ia *Processor) GetWidth() int {
	return ia.img.Bounds().Max.X
}

func (ia *Processor) GetHeight() int {
	return ia.img.Bounds().Max.Y
}

func (ia *Processor) GetExtension() string {
	return filepath.Ext(ia.filePath)
}

func (ia *Processor) GetAspectRatio() string {
	width := ia.GetWidth()
	height := ia.GetHeight()
	ratio := float64(width) / float64(height)

	commonRatios := []struct {
		Width, Height int
		RatioStr      string
	}{
		{16, 9, "16:9"},
		{9, 16, "9:16"},
		{21, 9, "21:9"},
		{4, 3, "4:3"},
		{3, 4, "3:4"},
		{1, 1, "1:1"},
	}

	var closestRatio string
	var minDifference float64 = math.MaxFloat64
	for _, r := range commonRatios {
		predefinedRatio := float64(r.Width) / float64(r.Height)
		difference := math.Abs(ratio - predefinedRatio)
		if difference < minDifference {
			minDifference = difference
			closestRatio = r.RatioStr
		}
	}

	return closestRatio
}

func (ia *Processor) GetSize() (int64, error) {
	fileInfo, err := os.Stat(ia.filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}

func (ia *Processor) GetMostFrequentColor(downSampleFactor int) string {
	colorFrequency := make(map[string]int)

	bounds := ia.img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for y := 0; y < height; y += downSampleFactor {
		for x := 0; x < width; x += downSampleFactor {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				c := ia.img.At(x, y)
				r, g, b, _ := c.RGBA()
				hexColor := rgbToHex(int(r>>8), int(g>>8), int(b>>8))

				mu.Lock()
				colorFrequency[hexColor]++
				mu.Unlock()
			}(x, y)
		}
	}

	wg.Wait()

	var mostFrequentColor string
	var maxCount int
	for color, count := range colorFrequency {
		if count > maxCount {
			maxCount = count
			mostFrequentColor = color
		}
	}

	return mostFrequentColor
}

func (ia *Processor) GenerateThumbnail(newWidth int) (image.Image, error) {
	bounds := ia.img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	aspectRatio := float64(width) / float64(height)
	newHeight := int(float64(newWidth) / aspectRatio)

	return imaging.Resize(ia.img, newWidth, newHeight, imaging.Lanczos), nil
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
