package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junioralcant/opportinities.git/schemas"
)

func CreateOpeningController(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createOpening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&createOpening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", createOpening)
}
