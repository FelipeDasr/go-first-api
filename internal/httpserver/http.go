package httpserver

import (
	"go-databases/internal/httpserver/router"

	"github.com/gin-gonic/gin"
)

// Start starts the HTTP server
func Start() {
	app := gin.Default();

	appRouter := router.AppHttpRouter{Router: app}
	appRouter.Init()

	app.Run(":3333")
}