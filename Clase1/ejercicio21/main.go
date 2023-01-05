package main

import (
	"ejemplos/ejercicio21/handlers"
	"encoding/json"
	"io"
	"os"

	"ejemplos/ejercicio21/globals"

	"github.com/gin-gonic/gin"
)

func main() {
	err := ReadFile()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	router := r.Group("/products")

	r.GET("/ping", handlers.Ping)
	router.GET("", handlers.Productos)
	router.GET("/:id", handlers.GetById)
	router.GET("/search", handlers.Search)

	r.Run()

}

func ReadFile() (err error) {
	file, err := os.Open("products1.json")
	if err != nil {
		return globals.ErrOpen
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return globals.ErrRead
	}
	err = json.Unmarshal(bytes, &globals.Products)
	if err != nil {
		return globals.ErrJson
	}
	return
}
