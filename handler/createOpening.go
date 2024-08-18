package handler

import (
	"github.com/ViniciusDSLima/streamming-golang/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(c *gin.Context) {

	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	err := request.Validate()

	if err != nil {
		logger.Errf("Validation Error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Link:     request.Link,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("Error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(c, "create-opening", &opening)
}
