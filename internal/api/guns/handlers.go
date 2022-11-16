package guns

import (
	Redis "go-api/internal/redis"
	Shared "go-api/internal/shared"
	"net/http"

	GunModels "go-api/internal/api/guns/models"
	GunSchemas "go-api/internal/api/guns/schemas"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetGuns(context *gin.Context) {
	hashes := []string{"guns:1", "guns:2", "guns:3", "guns:4"}
	var gunsModels = Redis.GetManyByHashes[GunModels.Gun](hashes)
	data := []any{}

	for _, model := range gunsModels {
		data = append(data, model.Data)
	}

	Shared.HandleResponse(context, http.StatusOK, data)
}

func PostGun(context *gin.Context) {
	var data []any
	var schema GunSchemas.GunPostSchema
	var model *GunModels.GunModel
	var created bool

	if error := context.ShouldBindWith(&schema, binding.JSON); error != nil {
		Shared.HandleRequestError(error, context)
		return
	}

	model = GunModels.NewGunModel(schema.GetGun())
	if created = Redis.CreateOne(model); !created {
		data = append(data, "Already exists")
		Shared.HandleErrorResponse(context, http.StatusConflict, data)
		return
	}
	data = append(data, GunSchemas.NewGunResponseSchema(model))
	Shared.HandleResponse(context, http.StatusAccepted, data)

}

func GetGunBySerialNumber(context *gin.Context) {
	data := []any{}
	serial_number := context.Param("serial_number")

	var hash = GunModels.GetGunHash(serial_number)
	var model = Redis.GetOne[GunModels.Gun](hash)

	if model == nil {
		Shared.HandleResponse(context, http.StatusNotFound, data)
	} else {
		var schema = GunSchemas.NewGunResponseSchema(model)
		data = append(data, schema)
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
