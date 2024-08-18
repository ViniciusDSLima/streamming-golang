package handler

import (
	"fmt"
	"github.com/ViniciusDSLima/streamming-golang/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deliting Opening with id: %s", id))
		return
	}

	sendSuccess(c, "delete+opening", &opening)
}
