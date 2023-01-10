package services

import (
	"ejemplos/ejercicio22/services/models"
	"errors"
	"fmt"
)

var (
	ErrAlreadyExist = errors.New("error: item already exist")
	ErrNotExist     = errors.New("error: item not exist")
	ErrOpen         = errors.New("\nError al abrir el archivo.")
	ErrRead         = errors.New("\nError al leer el archivo.")
	ErrJson         = errors.New("\nError al transformar el archivo Json.")
)

var Products = []models.Product{}
var lastID = 3

// read
func Get() []models.Product {
	return Products
}
func GetByID(id int) (models.Product, error) {
	for _, product := range Products {
		if product.Id == id {
			return product, nil
		}
	}
	return models.Product{}, ErrNotExist
}

func ExistCode(code string) bool {
	for _, w := range Products {
		if w.CodeValue == code {
			return true
		}
	}

	return false
}

func Create(name string, quantity int, code string, is_published bool, expiration string, price float64) (models.Product, error) {
	// validations
	if ExistCode(code) {
		return models.Product{}, fmt.Errorf("%w. %s", ErrAlreadyExist, "code value not unique")
	}

	lastID++
	product := models.Product{
		Id:          lastID,
		Name:        name,
		Quantity:    quantity,
		CodeValue:   code,
		IsPublished: is_published,
		Expiration:  expiration,
		Price:       price,
	}

	Products = append(Products, product)
	return product, nil
}
