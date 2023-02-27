package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"rest3/cmd/server/routes"
	"rest3/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// go get -u github.com/swaggo/swag/cmd/swag, go install github.com/swaggo/swag/cmd/swag@latest
// export PATH=$(go env GOPATH)/bin:$PATH,  swag init -g cmd/server/main.go
// http://localhost:8080/docs/index.html

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// instances
	if err := godotenv.Load("./cmd/server/.env"); err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	rt := routes.NewRouter(r)
	rt.SetRoutes()

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func ReadFile(prods *[]domain.Product) error {
	file, err := os.Open("products1.json")
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &prods)
	if err != nil {
		return err
	}
	return nil
}
