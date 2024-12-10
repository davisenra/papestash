package context

import (
	"fmt"
	"os"
	"strconv"

	"github.com/davisenra/papestash/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AppContext struct {
	HttpServer *gin.Engine
	Database   *database.Database
}

type AppConfig struct {
	DatabasePath string
	AppPort      int
}

func LoadAppConfig() (*AppConfig, error) {
	var missingVars []string

	getEnv := func(key string) string {
		value := os.Getenv(key)
		if value == "" {
			missingVars = append(missingVars, key)
		}
		return value
	}

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	appPortStr := getEnv("APP_PORT")

	var appPort int
	if appPortStr != "" {
		var err error
		appPort, err = strconv.Atoi(appPortStr)
		if err != nil {
			return nil, fmt.Errorf("invalid APP_PORT value: %v", err)
		}
	}

	config := &AppConfig{
		DatabasePath: getEnv("DB_PATH"),
		AppPort:      appPort,
	}

	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return config, nil
}
