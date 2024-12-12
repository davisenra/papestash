package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler func(*gin.Context) error

func Handle(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			log.Println("ERROR: " + err.Error())

			c.JSON(500, gin.H{
				"error": "Internal Server Error",
			})
		}
	}
}
