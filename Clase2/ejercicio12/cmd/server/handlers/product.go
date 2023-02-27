package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest3/internal/product"
	"strconv"

	//"rest2/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	ErrFormat = errors.New("Error: invalid format")
	ErrId     = errors.New("Error: invalid id")
	ErrQuery  = errors.New("Error: invalid query")
)

type ProductHandler struct {
	service product.Service
}

func NewProduct(service product.Service) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		prods, err := handler.service.Get()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, prods)
	}
}

func (handler *ProductHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ErrId.Error())
			return
		}
		prod, err := handler.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, prod)
	}
}

func (handler *ProductHandler) GetSearch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		priceGt := ctx.Query("price")
		price, err := strconv.ParseFloat(priceGt, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, ErrQuery.Error())
			return
		}
		prods, err := handler.service.GetSearch(price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, prods)
	}
}

func (handler *ProductHandler) Create() gin.HandlerFunc {
	type request struct {
		Name        string  `json:"name" validate:"required"`
		Quantity    int     `json:"quantity" validate:"required"`
		CodeValue   string  `json:"code_value" validate:"required"`
		IsPublished *bool   `json:"is_published"`
		Expiration  string  `json:"expiration" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
	}

	return func(ctx *gin.Context) {

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrFormat.Error())
			return
		}

		validate := validator.New()
		if err := validate.Struct(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}
		prod, err := handler.service.Create(req.Name, req.CodeValue, req.Expiration, req.Quantity, req.Price, *req.IsPublished)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, prod)
	}
}

func (handler *ProductHandler) Put() gin.HandlerFunc {
	type request struct {
		Name        string  `json:"name" validate:"required"`
		Quantity    int     `json:"quantity" validate:"required"`
		CodeValue   string  `json:"code_value" validate:"required"`
		IsPublished *bool   `json:"is_published"`
		Expiration  string  `json:"expiration" validate:"required"`
		Price       float64 `json:"price" validate:"required"`
	}

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrId.Error()})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrFormat.Error()})
			return
		}

		validate := validator.New()
		if err := validate.Struct(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		prod, err := handler.service.Update(int(id), req.Name, req.CodeValue, req.Expiration, req.Quantity, req.Price, *req.IsPublished)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, prod)
	}
}

func (handler *ProductHandler) Patch() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrId.Error()})
			return
		}

		prodOld, err := handler.service.GetById(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err = json.NewDecoder(ctx.Request.Body).Decode(&prodOld)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		prod, err := handler.service.PartialUpdate(id, prodOld)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, prod)
	}
}

func (handler *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		err = handler.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "product deleted"})
	}
}
