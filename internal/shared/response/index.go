package shared_response

import (
	"os"

	"github.com/gin-gonic/gin"
)

func sendResponse(context *gin.Context, status int, data any) {
	if os.Getenv("ENVIRONMENT") != "PRODUCTION" {
		context.IndentedJSON(status, data)
	} else {
		context.JSON(status, data)
	}
}
