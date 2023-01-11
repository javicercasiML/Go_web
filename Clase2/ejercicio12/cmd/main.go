package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"rest3/cmd/routes"
	"rest3/internal/domain"

	"github.com/gin-gonic/gin"
)

var (
	ErrOpen = errors.New("\nError al abrir el archivo.")
	ErrRead = errors.New("\nError al leer el archivo.")
	ErrJson = errors.New("\nError al transformar el archivo Json.")
)

func main() {
	// instances
	var db = []domain.Product{}
	err := ReadFile(&db)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	rt := routes.NewRouter(r, &db)
	rt.SetRoutes()

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func ReadFile(prods *[]domain.Product) error {
	file, err := os.Open("products1.json")
	if err != nil {
		return ErrOpen
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return ErrRead
	}
	err = json.Unmarshal(bytes, &prods)
	if err != nil {
		return ErrJson
	}
	return nil
}
