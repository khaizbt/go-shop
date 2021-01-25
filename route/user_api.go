package route

import (
	"goshop/config"
	"goshop/controller"
	"goshop/middleware"
	"goshop/service"

	"github.com/gin-gonic/gin"
)

var authService = config.NewServiceAuth()

func RouteUser(route *gin.Engine, service service.UserService) {

	userController := controller.NewUserController(service, authService)
	api := route.Group("/api/v1/user")
	api.POST("login", userController.Login)
	api.POST("update-account", middleware.AuthMiddlewareUser(authService, service, 0), userController.UpdateProfile)
	api.POST("create", middleware.AuthMiddlewareUser(authService, service, 1), userController.CreateUser)
	api.GET("list", middleware.AuthMiddlewareUser(authService, service, 4), userController.ListUser)
}
