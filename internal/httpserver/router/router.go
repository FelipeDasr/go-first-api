package router

import "github.com/gin-gonic/gin"

type AppHttpRouter struct {
	Router *gin.Engine
}

// Init initializes the APP HTTP router
func (r *AppHttpRouter) Init() {
	startProductRouter(r.Router)
	startCustomerRouter(r.Router)
	startOrderRouter(r.Router)
}