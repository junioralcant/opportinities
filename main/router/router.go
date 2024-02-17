package router

import (
	"github.com/gin-gonic/gin"
	"github.com/junioralcant/opportinities.git/infra/repositories"
)

func Initialize() {
	r := gin.Default()

	initializeRoutes(r)

	repositories.InitializeRepository()

	r.Run(":8080")
}
