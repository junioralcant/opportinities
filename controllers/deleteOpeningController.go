package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete World!"})
}
