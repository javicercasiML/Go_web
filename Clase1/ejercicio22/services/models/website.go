package models

import "errors"

type Product struct {
	Id          int     `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var (
	ErrOpen = errors.New("\nError al abrir el archivo.")
	ErrRead = errors.New("\nError al leer el archivo.")
	ErrJson = errors.New("\nError al transformar el archivo Json.")
)

var Products []Product
