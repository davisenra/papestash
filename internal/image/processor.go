package image

import (
	"fmt"
	"image"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
)

const DEFAULT_DOWNSAMPLE_FACTOR = 8

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

func (ia *Processor) Width() int {
	return ia.img.Bounds().Max.X
}

func (ia *Processor) Height() int {
	return ia.img.Bounds().Max.Y
}

func (ia *Processor) Extension() string {
	ext := filepath.Ext(ia.filePath)
	return strings.Replace(ext, ".", "", 1)
}

func (ia *Processor) AspectRatio() string {
	width := ia.Width()
	height := ia.Height()
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

func (ia *Processor) Size() int {
	fileInfo, _ := os.Stat(ia.filePath) // os.Stat might return an error, but hopefully it doesn't
	return int(fileInfo.Size())
}

func (ia *Processor) MostFrequentColor(downSampleFactor int) string {
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
