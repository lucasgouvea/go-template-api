package users

import (
	"fmt"
	"net/http"

	Shared "go-template-api/internal/shared"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/users", GetUsers)
	router.POST("/users", PostUser)
}

func GetUsers(c *gin.Context) {
	query := c.Request.URL.Query()
	length := len(query)
	fmt.Print(length)
	users := []User{{Name: "admin"}, {Name: "user_1"}}
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	user := User{}
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		httpError := Shared.GetHttpError(err)
		c.JSON(httpError.Status, httpError)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
