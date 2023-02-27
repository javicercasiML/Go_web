package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"rest3/cmd/server/routes"
	"rest3/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	ErrOpen = errors.New("\nError al abrir el archivo.")
	ErrRead = errors.New("\nError al leer el archivo.")
	ErrJson = errors.New("\nError al transformar el archivo Json.")
)

func main() {
	// instances
	if err := godotenv.Load("./cmd/server/.env"); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	rt := routes.NewRouter(r)
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
