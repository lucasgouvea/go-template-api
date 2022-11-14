package guns

import (
	"fmt"
	Redis "go-api/internal/redis"
	Shared "go-api/internal/shared"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetGuns(context *gin.Context) {
	hashes := []string{"guns:1", "guns:2", "guns:3", "guns:4"}
	var guns = Redis.GetMany[IGun](hashes)
	Shared.HandleResponse(context, http.StatusOK, guns)
}

func PostGun(context *gin.Context) {
	var newGun Shared.Model[IGun]

	error := context.ShouldBindWith(&newGun.Data, binding.JSON)

	if error != nil {
		Shared.HandleRequestError(error, context)
	} else {
		Guns = append(Guns, newGun)
		Shared.HandleResponse(context, http.StatusAccepted, newGun)
	}

}

func GetGunById(context *gin.Context) {
	id := context.Param("id")

	for _, gun := range Guns {
		if gun.Data.Id == id {
			Shared.HandleResponse(context, http.StatusOK, gun.Data)
			return
		}
	}

	Shared.HandleResponse(context, http.StatusNotFound, gin.H{"message": "not found"})
}

func GetAsync(context *gin.Context) {
	channel := make(chan int)
	go doSomethingAsync(channel)
	data := <-channel
	Shared.HandleResponse(context, http.StatusOK, data)

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
