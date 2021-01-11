package routers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/controllers"
	"wedding-invitation-go/middleware"
)

func InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(middleware.Cors())

	router.POST("/api/v1/admin/login", controllers.Login)

	return router
}
