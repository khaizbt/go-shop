package route

import (
	"github.com/gin-gonic/gin"
	"goshop/controller"
	"goshop/middleware"
	"goshop/repository"
	"goshop/service"
)

func ProductRoute(route *gin.Engine, services service.UserService) {
	productRepo := repository.NewProductCategory()
	productService := service.NewProductService(productRepo)
	productHandler := controller.ProductController(productService)
	api := route.Group("/api/v1/product")
	api.POST("create", middleware.AuthMiddlewareUser(authService, services, 5), productHandler.CreateProduct)
	api.GET("list-all", productHandler.ListProduct)
}
