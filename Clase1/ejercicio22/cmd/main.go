package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"ejemplos/ejercicio22/cmd/handlers"
	"ejemplos/ejercicio22/services"
)

func main() {
	err := ReadFile()
	if err != nil {
		panic(err)
	}

	//server
	r := gin.Default()

	//router
	router := r.Group("/products")
	router.GET("", handlers.Get)
	router.GET("/:id", handlers.GetByID)
	router.POST("", handlers.Create)

	// start
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func ReadFile() (err error) {
	file, err := os.Open("products1.json")
	if err != nil {
		return services.ErrOpen
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return services.ErrRead
	}
	err = json.Unmarshal(bytes, &services.Products)
	if err != nil {
		return services.ErrJson
	}
	return
}
