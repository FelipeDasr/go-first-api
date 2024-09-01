package httperror

import "github.com/gin-gonic/gin"

// HandleError handles the error
func HandleError(ctx *gin.Context, err *error) {
	ctx.JSON(400, gin.H{"error": (*err).Error()})
}