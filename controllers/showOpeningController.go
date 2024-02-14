package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Show World!"})
}
