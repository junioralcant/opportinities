package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List World!"})
}
