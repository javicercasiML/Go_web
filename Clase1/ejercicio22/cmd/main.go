package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"ejemplos/ejercicio22/cmd/handlers"
	"ejemplos/ejercicio22/services/models"
)

func main() {
	err := ReadFile()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router := r.Group("/products")

	//r.GET("/ping", handlers.Ping)
	router.GET("", handlers.Get)
	//router.GET("/:id", handlers.GetById)
	router.POST("", handlers.Create)

	r.Run()

}

func ReadFile() (err error) {
	file, err := os.Open("products1.json")
	if err != nil {
		return models.ErrOpen
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return models.ErrRead
	}
	err = json.Unmarshal(bytes, &models.Products)
	if err != nil {
		return models.ErrJson
	}
	return
}
