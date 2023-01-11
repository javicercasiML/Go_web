package product

import (
	"rest3/internal/domain"
)

type Service interface {
	Get() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	Create(name, code, expiration string, quantity int, price float64, is_publi bool) (domain.Product, error)
	GetSearch(price float64) ([]domain.Product, error)
	Update(id int, name, code, expiration string, quantity int, price float64, is_publi bool) (domain.Product, error)
	Delete(id int) error
	PartialUpdate(id int, prod domain.Product) (domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (service *service) Get() ([]domain.Product, error) {
	return service.repo.Get()
}

func (service *service) GetById(id int) (domain.Product, error) {
	return service.repo.GetById(id)
}

func (service *service) Create(name, code, expiration string, quantity int, price float64, is_publi bool) (domain.Product, error) {

	prod := domain.Product{
		Name:        name,
		Quantity:    quantity,
		CodeValue:   code,
		IsPublished: is_publi,
		Expiration:  expiration,
		Price:       price,
	}

	prod, err := service.repo.Create(prod)
	if err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}

func (service *service) GetSearch(price float64) ([]domain.Product, error) {
	prods, err := service.repo.GetSearch(price)
	if err != nil {
		return []domain.Product{}, err
	}
	return prods, nil
}

func (service *service) Update(id int, name, code, expiration string, quantity int, price float64, is_publi bool) (domain.Product, error) {

	prod := domain.Product{
		Id:          id,
		Name:        name,
		Quantity:    quantity,
		CodeValue:   code,
		IsPublished: is_publi,
		Expiration:  expiration,
		Price:       price,
	}

	prod, err := service.repo.Update(id, prod)
	if err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}

func (service *service) PartialUpdate(id int, prod domain.Product) (domain.Product, error) {

	prod, err := service.repo.Update(id, prod)
	if err != nil {
		return domain.Product{}, err
	}
	return prod, nil
}

func (service *service) Delete(id int) error {
	if err := service.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
