package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junioralcant/opportinities.git/schemas"
)

func ListOpeningController(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error in list openings: %v", err))
		return
	}

	sendSuccess(ctx, "list-opening", openings)
}
