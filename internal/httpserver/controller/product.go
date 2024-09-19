package controller

import (
	"errors"
	"go-databases/internal/httpserver/httperror"
	"go-databases/internal/service"
	"strconv"

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

func (pc *ProductController) GetProductById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		httperror.HandleError(ctx, errors.New("the id must be a number"))
		return
	}

	product, err := pc.ProductService.GetProductById(int32(id))

	if err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, product)
}

func (pc *ProductController) GetManyProducts(ctx *gin.Context) {
	var params service.FindManyProductsParams
	if err := ctx.ShouldBind(&params); err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	products, err := pc.ProductService.GetManyProducts(&params)

	if err != nil {
		httperror.HandleError(ctx, err)
		return
	}

	ctx.JSON(200, products)
}