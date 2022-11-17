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
	var schemas []GunSchemas.GunResponseSchema
	var models = *Redis.List[GunModels.Gun](GunModels.GunModelName)

	for _, model := range models {
		schemas = append(schemas, GunSchemas.NewGunResponseSchema(model))
	}

	Shared.HandleResponse(context, http.StatusOK, Shared.NewResponse(schemas))
}

func PostGun(context *gin.Context) {
	var data []any
	var schemas []GunSchemas.GunResponseSchema
	var postSchema GunSchemas.GunPostSchema
	var model *GunModels.GunModel

	if error := context.ShouldBindWith(&postSchema, binding.JSON); error != nil {
		Shared.HandleRequestError(error, context)
		return
	}

	model = GunModels.NewGunModel(postSchema.GetGun())

	if created := Redis.CreateOne(model); !created {
		data = append(data, "Already exists")
		Shared.HandleErrorResponse(context, http.StatusConflict, data)
		return
	}

	schemas = append(schemas, GunSchemas.NewGunResponseSchema(*model))
	Shared.HandleResponse(context, http.StatusAccepted, Shared.NewResponse(schemas))
}

func GetGunBySerialNumber(context *gin.Context) {
	var schemas []GunSchemas.GunResponseSchema
	serial_number := context.Param("serial_number")

	var hash = GunModels.GetGunHash(serial_number)
	var model = Redis.GetOne[GunModels.Gun](hash)

	if model == nil {
		Shared.HandleResponse(context, http.StatusNotFound, Shared.NewResponse(schemas))
	} else {
		schemas = append(schemas, GunSchemas.NewGunResponseSchema(*model))
		Shared.HandleResponse(context, http.StatusOK, Shared.NewResponse(schemas))
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
