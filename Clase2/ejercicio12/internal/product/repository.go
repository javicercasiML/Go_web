package product

import (
	"errors"
	"rest3/internal/domain"
)

type Repository interface {
	Get() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Create(prod domain.Product) (domain.Product, error)
	ExistCode(code string) bool
	GetSearch(price float64) ([]domain.Product, error)
	Update(id int, prod domain.Product) (domain.Product, error)
	Delete(id int) error
}

var (
	ErrNotFound      = errors.New("error: item not found")
	ErrDuplicateCode = errors.New("error: item duplicate")
)

type repository struct {
	db     *[]domain.Product
	lastID int
}

func NewRepository(db *[]domain.Product, lastID int) Repository {
	return &repository{db: db, lastID: lastID}
}

func (r *repository) Get() ([]domain.Product, error) {
	return *r.db, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	for _, prod := range *r.db {
		if prod.Id == id {
			return prod, nil
		}
	}
	return domain.Product{}, ErrNotFound
}

func (r *repository) Create(prod domain.Product) (domain.Product, error) {
	if r.ExistCode(prod.CodeValue) {
		return domain.Product{}, ErrDuplicateCode
	}
	r.lastID++
	prod.Id = r.lastID
	*r.db = append(*r.db, prod)
	return prod, nil
}

func (r *repository) GetSearch(price float64) ([]domain.Product, error) {
	var filtered []domain.Product
	for _, prod := range *r.db {
		if prod.Price > price {
			filtered = append(filtered, prod)
		}
	}
	if len(filtered) == 0 {
		return []domain.Product{}, ErrNotFound
	}
	return filtered, nil
}
func (r *repository) ExistCode(code string) (boolean bool) {
	for _, prod := range *r.db {
		if prod.CodeValue == code {
			return true
		}
	}
	return
}
func (r *repository) Update(id int, product domain.Product) (domain.Product, error) {
	for i, prod := range *r.db {
		if prod.Id == id {
			if r.ExistCode(product.CodeValue) && prod.CodeValue != product.CodeValue {
				return domain.Product{}, ErrDuplicateCode
			}
			(*r.db)[i] = product
			return product, nil
		}
	}
	return domain.Product{}, ErrNotFound
}

func (r *repository) Delete(id int) error {
	for i, prod := range *r.db {
		if prod.Id == id {
			(*r.db) = append((*r.db)[:i], (*r.db)[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
