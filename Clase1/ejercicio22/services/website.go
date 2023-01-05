package services

import (
	"ejemplos/ejercicio22/services/models"
	"errors"
	"fmt"
)

var (
	ErrAlreadyExist = errors.New("error: item already exist")
)

var lastID = 3

// read
func Get() []models.Product {
	return models.Products
}
func GetByID() {

}
func ExistCode(code string) bool {
	for _, w := range models.Products {
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

	models.Products = append(models.Products, product)
	return product, nil
}
