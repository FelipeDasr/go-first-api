package router

import (
	"github.com/gin-gonic/gin"
)

func startCustomerRouter(router *gin.Engine) {
	customerRouter := router.Group("/customers")
}