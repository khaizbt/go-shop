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
	categgoryHandler := controller.CategoryController(categoryService)
	api := route.Group("/api/v1/category")
	api.POST("create", middleware.AuthMiddlewareUser(authService, services, 5), categgoryHandler.CreateCategory)
	api.PUT("update/:id", middleware.AuthMiddlewareUser(authService, services, 6), categgoryHandler.UpdateCategory)

}
