package router

import (
	"go-databases/internal/httpserver/controller"
	"go-databases/internal/service"

	"github.com/gin-gonic/gin"
)

func startProductRouter(router *gin.Engine) {
	productService := service.NewProductService()
	productController := controller.NewProductController(productService)

	productRouter := router.Group("/products")

	productRouter.POST("/", productController.CreateProduct)
	productRouter.GET("/", productController.GetManyProducts)
	productRouter.GET("/:id", productController.GetProductById)
}