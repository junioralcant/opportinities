package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.GET("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Get World!"})
	})

	v1.POST("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Post World!"})
	})

	v1.DELETE("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Delete World!"})
	})

	v1.PUT("/opening", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Put World!"})
	})

	v1.GET("/openings", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Get all World!"})
	})
}
