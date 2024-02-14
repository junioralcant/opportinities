package router

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/opportinities.git/controllers"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.GET("/opening", controllers.ShowOpeningController)

	v1.POST("/opening", controllers.CreateOpeningController)

	v1.DELETE("/opening", controllers.DeleteOpeningController)

	v1.PUT("/opening", controllers.UpdateOpeningController)

	v1.GET("/openings", controllers.ListOpeningController)
}
