package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	// Corremos nuestro servidor sobre el puerto 8080
	if err := router.Run(); err != nil {
		panic(err)
	}

}
