package utils

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func SluggifyFile(file string) string {
	originalName := strings.TrimSuffix(file, filepath.Ext(file))
	slug := strings.ReplaceAll(strings.ToLower(originalName), " ", "-")
	uniqueID := time.Now().UnixNano()
	return fmt.Sprintf("%s-%d%s", slug, uniqueID, filepath.Ext(file))
}
