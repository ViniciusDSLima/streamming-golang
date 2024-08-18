package router

import "github.com/gin-gonic/gin"

// Initialize start the project
func Initialize() {

	router := gin.Default()

	initializeRoutes(router)

	router.Run(":8080")
}
