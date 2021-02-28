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
	api.GET("list", middleware.AuthMiddlewareUser(authService, services, 0), productHandler.ListProductUser)
	api.PUT("update/:id", middleware.AuthMiddlewareUser(authService, services, 5), productHandler.UpdateProduct)
	api.DELETE("delete/:id", middleware.AuthMiddlewareUser(authService, services, 5), productHandler.DeleteProduct)
}
