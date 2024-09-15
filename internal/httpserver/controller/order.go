package controller

import (
	"go-databases/internal/httpserver/httperror"
	"go-databases/internal/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService *service.OrderService
}

// NewCustomerController creates a new CustomerController
func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{OrderService: orderService}
}

// CreateOrder creates a new order
func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var data service.CreateOrderData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	result, err := oc.OrderService.CreateOrder(&data)

	if err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	ctx.JSON(201, result)
}