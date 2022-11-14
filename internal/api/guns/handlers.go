package guns

import (
	Redis "go-api/internal/redis"
	Shared "go-api/internal/shared"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetGuns(context *gin.Context) {
	hashes := []string{"guns:1", "guns:2", "guns:3", "guns:4"}
	var gunsModels = Redis.GetMany[Gun](hashes)
	data := []any{}

	for _, gunModel := range gunsModels {
		data = append(data, gunModel.Data)
	}

	Shared.HandleResponse(context, http.StatusOK, data)
}

func PostGun(context *gin.Context) {
	var newGun Redis.Model[Gun]
	data := []any{}

	error := context.ShouldBindWith(&newGun.Data, binding.JSON)

	data = append(data, newGun.Data)

	if error != nil {
		Shared.HandleRequestError(error, context)
	} else {
		Shared.HandleResponse(context, http.StatusAccepted, data)
	}

}

func GetGunById(context *gin.Context) {
	data := []any{}
	id := context.Param("id")

	var gun = Redis.GetOne[Gun]("guns:" + id)

	if gun == nil {
		Shared.HandleResponse(context, http.StatusNotFound, data)
	} else {
		data = append(data, gun.Data)
		Shared.HandleResponse(context, http.StatusOK, data)
	}

}

/* func GetAsync(context *gin.Context) {
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
} */
