package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"message": "TO BE IMPLEMENTED",
	})

	return nil
}
