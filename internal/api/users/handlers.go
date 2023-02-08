package users

import (
	"net/http"

	Shared "go-template-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/users", GetUsers)
	router.POST("/users", PostUser)
	router.PATCH("/users/:id", PatchUser)
}

func GetUsers(ctx *gin.Context) {
	var err error
	var params Shared.Params
	var users []User

	query := ctx.Request.URL.Query()

	if params, err = Shared.ParseQuery(query); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if users, err = listUsers(params); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func PostUser(ctx *gin.Context) {
	schema := UserPostSchema{}
	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	user := schema.parse()

	if err := createUser(user); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func PatchUser(ctx *gin.Context) {
	var user *User
	var err error

	schema := UserPatchSchema{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindWith(&schema, binding.JSON); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if user, err = schema.parse(id); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	if err := updateUser(user); err != nil {
		Shared.HandleErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
