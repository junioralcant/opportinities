package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Post World!"})
}
