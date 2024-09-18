package router

import (
	"go-databases/internal/httpserver/controller"
	"go-databases/internal/service"

	"github.com/gin-gonic/gin"
)

func startOrderRouter(router *gin.Engine) {
	orderServices := service.NewOrderService()
	ordersController := controller.NewOrderController(orderServices)

	orderRouter := router.Group("/orders")

	orderRouter.POST("/", ordersController.CreateOrder)
	orderRouter.GET("/:id", ordersController.GetOrderById)
}