package guns

import (
	Redis "go-api/internal/redis"
	Shared "go-api/internal/shared"
	Response "go-api/internal/shared/response"
	"net/http"

	GunModels "go-api/internal/api/guns/models"
	GunSchemas "go-api/internal/api/guns/schemas"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetGuns(context *gin.Context) {
	var schemas = []GunSchemas.GunResponseSchema{}
	var query Shared.Query

	if err := context.ShouldBindWith(&query, binding.Query); err != nil {
		Shared.HandleRequestError(err, context)
		return
	}

	var models = *Redis.List[GunModels.Gun](GunModels.GunModelName, query)

	for _, model := range models {
		schemas = append(schemas, GunSchemas.NewGunResponseSchema(model))
	}

	Response.Send(context, http.StatusOK, Response.New(schemas))
}

func PostGun(context *gin.Context) {
	var schemas []GunSchemas.GunResponseSchema
	var postSchema GunSchemas.GunPostSchema
	var model *GunModels.GunModel

	if error := context.ShouldBindWith(&postSchema, binding.JSON); error != nil {
		Shared.HandleRequestError(error, context)
		return
	}

	model = GunModels.NewGunModel(postSchema.GetGun())

	if created := Redis.CreateOne(model); !created {
		Response.SendError(context, http.StatusConflict, Response.NewError([]string{"Already exists"}))
		return
	}

	schemas = append(schemas, GunSchemas.NewGunResponseSchema(*model))
	Response.Send(context, http.StatusAccepted, Response.New(schemas))
}

func GetGunBySerialNumber(context *gin.Context) {
	schemas := []GunSchemas.GunResponseSchema{}
	serial_number := context.Param("serial_number")

	var hash = GunModels.GetGunHash(serial_number)
	var model = Redis.GetOne[GunModels.Gun](hash)

	if model == nil {
		Response.Send(context, http.StatusNotFound, Response.New(schemas))
	} else {
		schemas = append(schemas, GunSchemas.NewGunResponseSchema(*model))
		Response.Send(context, http.StatusOK, Response.New(schemas))
	}

}

/* func GetAsync(context *gin.Context) {
	channel := make(chan int)
	go doSomethingAsync(channel)
	data := <-channel
	Response.Send(context, http.StatusOK, data)
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
