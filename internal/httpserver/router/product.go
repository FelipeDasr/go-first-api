package router

import (
	"github.com/gin-gonic/gin"
)

func startProductRouter(router *gin.Engine) {
	productRouter := router.Group("/products")
}