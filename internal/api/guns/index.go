package guns

import "github.com/gin-gonic/gin"

func Start(router *gin.RouterGroup) {
	router.GET("/guns", GetGuns)
	router.POST("/guns", PostGun)
	router.GET("/guns/:serial_number", GetGunBySerialNumber)
	//router.GET("/async", GetAsync)
}
