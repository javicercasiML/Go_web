package store

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"rest3/internal/domain"
)

type Store interface {
	loadProducts() ([]domain.Product, error)
	GetProducts() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Post(prod domain.Product) (domain.Product, error)
	Delete(id int) error
	Update(prodc domain.Product) error
}

var (
	ErrOpen     = errors.New("\nError al abrir el archivo.")
	ErrRead     = errors.New("\nError al leer el archivo.")
	ErrJson     = errors.New("json data invalid")
	ErrNotFound = errors.New("product not found")
)

type jsonStore struct {
	file string
}

func NewStorage(file string) Store {
	return &jsonStore{file}
}

func (store *jsonStore) loadProducts() ([]domain.Product, error) {
	var prods []domain.Product
	file, err := os.Open(store.file)
	if err != nil {
		return nil, ErrOpen
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, ErrRead
	}
	err = json.Unmarshal(bytes, &prods)
	if err != nil {
		return nil, ErrJson
	}
	return prods, nil
}

func (store *jsonStore) GetProducts() ([]domain.Product, error) {
	prods, err := store.loadProducts()
	if err != nil {
		return nil, err
	}
	return prods, nil
}

func (store *jsonStore) GetById(id int) (domain.Product, error) {
	prods, err := store.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}

	for _, prod := range prods {
		if prod.Id == id {
			return prod, nil
		}
	}
	return domain.Product{}, ErrNotFound
}

func (store *jsonStore) Post(prod domain.Product) (domain.Product, error) {
	prods, err := store.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	prod.Id = len(prods) + 1
	prods = append(prods, prod)
	products, err := json.Marshal(prods)
	if err != nil {
		return domain.Product{}, ErrJson
	}
	if err := os.WriteFile(store.file, products, 0644); err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}

func (store *jsonStore) Delete(id int) error {
	prods, err := store.loadProducts()
	if err != nil {
		return err
	}
	for i, prod := range prods {
		if prod.Id == id {
			prods = append(prods[:i], prods[i+1:]...)
			products, err := json.Marshal(prods)
			if err != nil {
				return ErrJson
			}
			if err := os.WriteFile(store.file, products, 0644); err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNotFound
}

func (store *jsonStore) Update(prodc domain.Product) error {
	prods, err := store.loadProducts()
	if err != nil {
		return err
	}
	for i, prod := range prods {
		if prod.Id == prodc.Id {
			prods[i] = prodc
			products, err := json.Marshal(prods)
			if err != nil {
				return ErrJson
			}
			if err := os.WriteFile(store.file, products, 0644); err != nil {
				return err
			}
			return nil
		}
	}
	return ErrNotFound
}
