package product

import (
	"errors"
	"rest3/internal/domain"
	"rest3/pkg/store"
)

type Repository interface {
	Get() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Create(prod domain.Product) (domain.Product, error)
	ExistCode(product domain.Product) (boolean bool)
	GetSearch(price float64) ([]domain.Product, error)
	Update(id int, prod domain.Product) (domain.Product, error)
	Delete(id int) error
}

var (
	ErrNotFound      = errors.New("error: item not found")
	ErrDuplicateCode = errors.New("error: item duplicate")
)

type repository struct {
	storage store.Store
}

func NewRepository(storage store.Store) Repository {
	return &repository{storage: storage}
}

func (r *repository) Get() ([]domain.Product, error) {
	prods, err := r.storage.GetProducts()
	if err != nil {
		return []domain.Product{}, err
	}
	return prods, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	prod, err := r.storage.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}

func (r *repository) Create(prod domain.Product) (domain.Product, error) {
	if r.ExistCode(prod) {
		return domain.Product{}, ErrDuplicateCode
	}
	prod, err := r.storage.Post(prod)
	if err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}

func (r *repository) GetSearch(price float64) ([]domain.Product, error) {
	var filtered []domain.Product
	prods, err := r.Get()
	if err != nil {
		return []domain.Product{}, nil
	}

	for _, prod := range prods {
		if prod.Price > price {
			filtered = append(filtered, prod)
		}
	}
	if len(filtered) == 0 {
		return []domain.Product{}, ErrNotFound
	}
	return filtered, nil
}
func (r *repository) ExistCode(product domain.Product) (boolean bool) {
	prods, _ := r.Get()

	for _, prod := range prods {
		if prod.CodeValue == product.CodeValue && prod.Id != product.Id {
			return true
		}
	}
	return
}
func (r *repository) Update(id int, product domain.Product) (domain.Product, error) {

	if r.ExistCode(product) {
		return domain.Product{}, ErrDuplicateCode
	}
	if err := r.storage.Update(product); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Delete(id int) error {
	if err := r.storage.Delete(id); err != nil {
		return ErrNotFound
	}
	return nil
}
