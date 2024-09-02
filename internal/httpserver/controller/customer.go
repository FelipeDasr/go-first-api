package controller

import (
	"errors"
	"go-databases/internal/httpserver/httperror"
	"go-databases/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CustomerService *service.CustomerServices
}

// NewCustomerController creates a new CustomerController
func NewCustomerController(customerService *service.CustomerServices) *CustomerController {
	return &CustomerController{CustomerService: customerService}
}

func (cc *CustomerController) CreateCustomer(ctx *gin.Context) {
	var data service.CreateCustomerData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	result, err := cc.CustomerService.CreateCustomer(&data)

	if err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	ctx.JSON(201, result)
}

func (cc *CustomerController) GetCustomerById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		httperror.HandleError(ctx, errors.New("invalid ID"))
		return
	}

	result, err := cc.CustomerService.GetCustomerById(int32(id))

	if err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, result)
}