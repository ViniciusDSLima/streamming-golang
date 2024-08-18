package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Get Openiing",
			})
		})

		v1.POST("/create-opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Get Openiing",
			})
		})

		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Get Openiing",
			})
		})

		v1.PATCH("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Get Openiing",
			})
		})

		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Get Openiings",
			})
		})
	}
}
