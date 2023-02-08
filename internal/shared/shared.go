package shared

import "github.com/gin-gonic/gin"

func HandleErr(ctx *gin.Context, err error) {
	httpError := GetHttpError(err)
	ctx.JSON(httpError.Status, httpError)
}
