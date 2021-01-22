package route

import (
	"goshop/config"
	"goshop/controller"
	"goshop/middleware"
	"goshop/service"

	"github.com/gin-gonic/gin"
)

func RouteUser(route *gin.Engine, service service.UserService) {
	authService := config.NewServiceAuth()
	userController := controller.NewUserController(service, authService)
	api := route.Group("/api/v1/user")
	api.POST("login", userController.Login)
	api.POST("update-account", middleware.AuthMiddlewareUser(authService, service, 0), userController.UpdateProfile)
	api.POST("create", middleware.AuthMiddlewareUser(authService, service, 1), userController.CreateUser)
}
