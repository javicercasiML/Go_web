package handlers

import (
	"ejemplos/ejercicio22/pkg/response"
	"ejemplos/ejercicio22/services"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const SECRET_KEY = "ABC"

var (
	ErrUnauthorized = errors.New("error: invalid token")
	ErrDate         = errors.New("Error: invalid date")
)

func Get(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token != SECRET_KEY {
		ctx.JSON(http.StatusUnauthorized, response.Err(ErrUnauthorized))
		return
	}

	// request

	// process
	products := services.Get()

	// response
	ctx.JSON(http.StatusOK, response.Ok("succeed to get products", products))
}

type request struct {
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished *bool   `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

func Create(ctx *gin.Context) {
	// request
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}
	aux := strings.Split(req.Expiration, "/")
	_, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", aux[2], aux[1], aux[0]))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(ErrDate))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	// process
	product, err := services.Create(req.Name, req.Quantity, req.CodeValue, *req.IsPublished, req.Expiration, req.Price)
	if err != nil {
		if errors.Is(err, services.ErrAlreadyExist) {
			ctx.JSON(http.StatusConflict, response.Err(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	// response
	ctx.JSON(http.StatusCreated, response.Ok("suceed to create product", product))
}
