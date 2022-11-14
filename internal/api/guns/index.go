package guns

import "github.com/gin-gonic/gin"

func Start(router *gin.RouterGroup) {
	router.GET("/guns", GetGuns)
	router.POST("/guns", PostGun)
	router.GET("/guns/:id", GetGunById)
	//router.GET("/async", GetAsync)
}
