package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"rest3/internal/product"
	"rest3/pkg/web"
	"strconv"

	//"rest2/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	ErrFormat = errors.New("invalid format")
	ErrId     = errors.New("invalid id")
	ErrQuery  = errors.New("invalid query")
	ErrToken  = errors.New("invalid Token")
)

type ProductHandler struct {
	service product.Service
}

func NewProduct(service product.Service) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("TOKEN")
		if token != os.Getenv("TOKEN") {
			web.Failed(ctx, http.StatusUnauthorized, ErrToken)
			return
		}
		prods, err := handler.service.Get()
		if err != nil {
			web.Failed(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusOK, prods)
	}
}

func (handler *ProductHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, ErrId)
			return
		}
		prod, err := handler.service.GetById(id)
		if err != nil {
			web.Failed(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, prod)
	}
}

func (handler *ProductHandler) GetSearch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		priceGt := ctx.Query("price")
		price, err := strconv.ParseFloat(priceGt, 64)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, ErrQuery)
			return
		}
		prods, err := handler.service.GetSearch(price)
		if err != nil {
			web.Failed(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, prods)
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

		token := ctx.GetHeader("TOKEN")
		if token != os.Getenv("TOKEN") {
			web.Failed(ctx, http.StatusUnauthorized, ErrToken)
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Failed(ctx, http.StatusBadRequest, ErrFormat)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&req); err != nil {
			web.Failed(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		prod, err := handler.service.Create(req.Name, req.CodeValue, req.Expiration, req.Quantity, req.Price, *req.IsPublished)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, err)
			return
		}

		web.Success(ctx, http.StatusOK, prod)
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

		token := ctx.GetHeader("TOKEN")
		if token != os.Getenv("TOKEN") {
			web.Failed(ctx, http.StatusUnauthorized, ErrToken)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ErrId.Error()})
			return
		}

		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Failed(ctx, http.StatusBadRequest, ErrFormat)
			return
		}

		validate := validator.New()
		if err := validate.Struct(&req); err != nil {
			web.Failed(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		prod, err := handler.service.Update(int(id), req.Name, req.CodeValue, req.Expiration, req.Quantity, req.Price, *req.IsPublished)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, err)
			return
		}

		web.Success(ctx, http.StatusOK, prod)
	}
}

func (handler *ProductHandler) Patch() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader("TOKEN")
		if token != os.Getenv("TOKEN") {
			web.Failed(ctx, http.StatusUnauthorized, ErrToken)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, ErrId)
			return
		}

		prodOld, err := handler.service.GetById(id)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, err)
			return
		}

		err = json.NewDecoder(ctx.Request.Body).Decode(&prodOld)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, err)
			return
		}

		prod, err := handler.service.PartialUpdate(id, prodOld)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, err)
			return
		}

		web.Success(ctx, http.StatusOK, prod)
	}
}

func (handler *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("TOKEN")
		if token != os.Getenv("TOKEN") {
			web.Failed(ctx, http.StatusUnauthorized, ErrToken)
			return
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, ErrId)
			return
		}
		err = handler.service.Delete(id)
		if err != nil {
			web.Failed(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, "Product deleted")
	}
}
