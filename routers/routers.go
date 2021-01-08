package routers

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/middleware"
)

func InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(middleware.Cors())

	return router
}
