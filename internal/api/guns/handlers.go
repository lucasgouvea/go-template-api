package guns

import (
	"fmt"
	Redis "go-api/internal/redis"
	Shared "go-api/internal/shared"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetGuns(context *gin.Context) {
	hashes := []string{"guns:1", "guns:2", "guns:3", "guns:4"}
	var guns = Redis.GetMany[IGun](hashes)
	context.JSON(http.StatusOK, guns)
}

func PostGun(context *gin.Context) {
	var newGun Shared.Model[IGun]

	if err := context.BindJSON(&newGun); err != nil {
		return
	}

	Guns = append(Guns, newGun)
	context.IndentedJSON(http.StatusCreated, newGun)
}

func GetGunById(context *gin.Context) {
	id := context.Param("id")

	for _, gun := range Guns {
		if gun.Data.Id == id {
			context.JSON(http.StatusOK, gun.Data)
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func GetAsync(context *gin.Context) {
	channel := make(chan int)
	go doSomethingAsync(channel)
	data := <-channel
	context.JSON(http.StatusOK, data)

}

func doSomethingAsync(channel chan int) chan int {
	fmt.Println("Warming up ...")
	go func() {
		time.Sleep(3 * time.Second)
		channel <- 1
		fmt.Println("Done ...")
	}()
	return channel
}
