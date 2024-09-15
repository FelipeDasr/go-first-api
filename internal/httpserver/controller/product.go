package controller

import (
	"go-databases/internal/httpserver/httperror"
	"go-databases/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var data service.CreateProductData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	result, err := pc.ProductService.CreateProduct(&data)

	if err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	ctx.JSON(201, result)
}