package context

import (
	"github.com/davisenra/papestash/internal/database"
	"github.com/gin-gonic/gin"
)

type AppContext struct {
	HttpServer *gin.Engine
	Database   *database.Database
}
