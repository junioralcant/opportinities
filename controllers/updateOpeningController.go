package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update World!"})
}
