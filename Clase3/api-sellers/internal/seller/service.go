package seller

import (
	"context"
	"errors"

	"grabaciones/functional_test/internal/domain"
)

// Errors
var (
	ErrNotFound      = errors.New("seller not found")
	ErrAlreadyExists = errors.New("seller already exists")
	ErrRepository    = errors.New("error in repository")

	EmptySeller = domain.Seller{}
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Create(ctx context.Context, s domain.Seller) (domain.Seller, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	sellers, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return sellers, nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Seller, error) {
	seller, err := s.repo.Get(ctx, id)

	// if errors.Is(err, ErrNotFound) {
	// 	return EmptySeller, ErrNotFound
	// }

	if err != nil {
		return EmptySeller, err
	}

	return seller, nil
}

func (s *service) Create(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	exist := s.repo.Exists(ctx, seller.CID)

	if exist {
		return EmptySeller, ErrAlreadyExists
	}

	id, err := s.repo.Save(ctx, seller)
	if err != nil {
		return EmptySeller, err
	}

	seller.ID = id

	return seller, nil
}
