package handlers

import (
	"ejemplos/ejercicio21/globals"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func Productos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, globals.Response{Message: "Succeed", Data: globals.Products})
}

func GetById(ctx *gin.Context) {
	// request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, globals.Response{Message: "Failed to parse number", Data: nil})
		return
	}

	// process
	var filtered []globals.Product
	for _, p := range globals.Products {
		if p.Id == id {
			filtered = append(filtered, p)
		}
	}

	// response
	ctx.JSON(http.StatusOK, globals.Response{Message: "succeed", Data: filtered})
}

func Search(ctx *gin.Context) {
	// request
	query, err := strconv.ParseFloat(ctx.Query("price"), 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, globals.Response{Message: "failed to parse number", Data: nil})
		return
	}

	// process
	var filtered []globals.Product
	for _, p := range globals.Products {
		if p.Price > query {
			filtered = append(filtered, p)
		}
	}

	// response
	ctx.JSON(http.StatusOK, globals.Response{Message: "succeed", Data: filtered})
}
