package shared

import (
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Limit  int
	Offset int
}

const DEFAULT_LIMIT = 10
const DEFAULT_OFFSET = 0

func HandleErr(ctx *gin.Context, err error) {
	httpError := GetHttpError(err)
	ctx.JSON(httpError.Status, httpError)
}

func ParseQuery(query url.Values) (Params, error) {
	var params Params
	var err error

	limit := query.Get("limit")
	offset := query.Get("offset")

	if len(limit) > 0 {
		params.Limit, err = strconv.Atoi(limit)
	} else {
		params.Limit = DEFAULT_LIMIT
	}

	if len(offset) > 0 {
		params.Offset, err = strconv.Atoi(offset)
	} else {
		params.Offset = DEFAULT_OFFSET
	}

	return params, err
}
