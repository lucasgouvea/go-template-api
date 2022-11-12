package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/guns", getGuns)
	router.POST("/guns", postGun)
	router.GET("/guns/:id", getGunById)
	router.GET("/async", getAsync)
	router.Run("localhost:8080")
}

type IGun struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var guns = []IGun{
	{Id: "1", Name: "AK-47", Price: 2000},
	{Id: "2", Name: "Glock", Price: 500},
	{Id: "3", Name: "MP5", Price: 1150},
}

func getGuns(context *gin.Context) {
	context.JSON(http.StatusOK, guns)
}

func postGun(context *gin.Context) {
	var newGun IGun

	if err := context.BindJSON(&newGun); err != nil {
		return
	}

	guns = append(guns, newGun)
	context.IndentedJSON(http.StatusCreated, newGun)
}

func getGunById(context *gin.Context) {
	id := context.Param("id")

	for _, gun := range guns {
		if gun.Id == id {
			context.JSON(http.StatusOK, gun)
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"message": "not found"})
}

func getAsync(context *gin.Context) {
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
