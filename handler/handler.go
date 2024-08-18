package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST Openiing",
	})
}

func ShowOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Openiing",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE Openiing",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Openiing",
	})
}

func ListOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List Openings",
	})
}
