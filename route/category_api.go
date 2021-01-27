package route

import (
	"github.com/gin-gonic/gin"
	"goshop/controller"
	"goshop/middleware"
	"goshop/repository"
	"goshop/service"
)

func CategoryRoute(route *gin.Engine, services service.UserService) {
	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := controller.CategoryController(categoryService)
	api := route.Group("/api/v1/category")
	api.POST("create", middleware.AuthMiddlewareUser(authService, services, 5), categoryHandler.CreateCategory)
	api.PUT("update/:id", middleware.AuthMiddlewareUser(authService, services, 6), categoryHandler.UpdateCategory)
	api.GET("list", middleware.AuthMiddlewareUser(authService, services, 7), categoryHandler.ListCategory)
	api.DELETE("delete/:id", middleware.AuthMiddlewareUser(authService, services, 8), categoryHandler.DeleteCategory)
}
