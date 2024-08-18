package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List Openings",
	})
}
