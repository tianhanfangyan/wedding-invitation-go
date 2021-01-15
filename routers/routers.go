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

	adminRouter := router.Group("/api/v1/admin")
	adminRouter.Use(middleware.JWTAuth())

	adminRouter.POST("/userinfo", controllers.AddNewComerInfo)
	adminRouter.DELETE("/userinfo", controllers.DeleteNewComerInfo)
	adminRouter.GET("/userinfo", controllers.GetNewComerInfo)
	adminRouter.PUT("/userinfo", controllers.ModifyNewComerInfo)

	apiRouter := router.Group("/api/v1/user")
	apiRouter.Use(middleware.JWTAuth())

	apiRouter.Static("/assets", "./assets")
	apiRouter.GET("/userinfo/", controllers.GetNewComerInfoByUserId)
	apiRouter.GET("/hotel", controllers.GetHotelLocationByUserId)

	return router
}
