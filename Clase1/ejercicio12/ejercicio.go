package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func main() {
	router := gin.Default()

	router.POST("/saludo", func(ctx *gin.Context) {
		var person Persona
		if err := ctx.BindJSON(&person); err != nil {
			ctx.String(http.StatusInternalServerError, "Bad format")
		} else {
			message := fmt.Sprintf("Hola %s %s", person.Nombre, person.Apellido)
			ctx.String(http.StatusOK, message)
		}
	})

	if err := router.Run(); err != nil {
		panic(err)
	}

}
