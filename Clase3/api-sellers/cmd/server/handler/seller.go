package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"grabaciones/functional_test/internal/domain"
	"grabaciones/functional_test/internal/seller"
	"grabaciones/functional_test/pkg/web"

	"github.com/gin-gonic/gin"
)

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idConv, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		seller, err := s.sellerService.Get(c, idConv)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, seller)
	}
}

func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Seller

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		validity := Validation(req)
		if validity != "" {
			web.Error(c, http.StatusUnprocessableEntity, validity)
			return
		}

		seller, err := s.sellerService.Create(c, req)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, seller)

	}
}

func Validation(req interface{}) string {
	reqValue := reflect.ValueOf(req)
	for i := 0; i < reqValue.NumField(); i++ {
		value := reqValue.Field(i).Interface()
		tipe := reflect.TypeOf(value).Kind()
		if fmt.Sprint(tipe) == "string" {
			if value == "" {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "int" && reqValue.Type().Field(i).Name != "ID" {
			if value.(int) <= 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "float64" {
			if value.(float64) == 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "boolean" {
			if !value.(bool) {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		}
	}
	return ""
}
