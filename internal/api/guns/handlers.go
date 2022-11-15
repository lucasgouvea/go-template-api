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
	var gunsModels = Redis.GetManyByHashes[Gun](hashes)
	data := []any{}

	for _, gunModel := range gunsModels {
		data = append(data, gunModel.Data)
	}

	Shared.HandleResponse(context, http.StatusOK, data)
}

func PostGun(context *gin.Context) {
	var data []any
	var gun Gun
	var gunModel *Redis.Model[Gun]

	if error := context.ShouldBindWith(&gun, binding.JSON); error != nil {
		Shared.HandleRequestError(error, context)
		return
	}

	gunModel = NewGunModel(gun)
	gunModel = Redis.CreateOne(gunModel)
	data = append(data, gunModel)

	Shared.HandleResponse(context, http.StatusAccepted, data)

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
