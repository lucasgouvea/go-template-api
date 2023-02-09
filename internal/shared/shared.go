package shared

import (
	"net/url"
	"strconv"
	"unicode"
	"unicode/utf8"

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

// https://github.com/fatih/camelcase/blob/master/camelcase.go
func SplitCamelCase(src string) (entries []string) {
	// don't split invalid utf8
	if !utf8.ValidString(src) {
		return []string{src}
	}
	entries = []string{}
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}
	// handle upper case -> lower case sequences, e.g.
	// "PDFL", "oader" -> "PDF", "Loader"
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}
	return
}
