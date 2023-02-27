package mocks

import (
	"context"
	"grabaciones/functional_test/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (domain.Seller, error)
	Save(ctx context.Context, s domain.Seller) (int, error)
	Exists(ctx context.Context, cid int) bool
}

type DbMock struct {
	SellerMocked []domain.Seller
	Err          error
	Exist        bool
}

type repoStub struct {
	db *DbMock
}

func NewRepositorySeller(dbM *DbMock) Repository {
	return &repoStub{dbM}
}

var CtxMock = context.Background()

func (s *repoStub) GetAll(ctx context.Context) ([]domain.Seller, error) {
	if s.db.Err != nil {
		return nil, s.db.Err
	}

	return s.db.SellerMocked, nil
}

func (s *repoStub) Get(ctx context.Context, id int) (domain.Seller, error) {

	if s.db.Err != nil {
		return domain.Seller{}, s.db.Err
	}

	return domain.Seller{}, nil
}

func (s *repoStub) Save(ctx context.Context, se domain.Seller) (int, error) {
	if s.db.Err != nil {
		return 0, s.db.Err
	}

	return se.ID, nil
}

func (s *repoStub) Exists(ctx context.Context, cid int) bool {
	return s.db.Exist
}
