package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/youngmikedev/warehouse/internal/domain"
	"github.com/youngmikedev/warehouse/internal/repository"
)

type ProductService struct {
	repo *repository.Repositories
	log  *zerolog.Logger
}

func NewProductService(repo *repository.Repositories, logger *zerolog.Logger) *ProductService {
	sl := logger.With().Str("service", "product").Logger()
	return &ProductService{
		repo: repo,
		log:  &sl,
	}
}

func (s *ProductService) Create(ctx context.Context, uid int, product domain.Product) (id int, err error) {
	id, err = s.repo.Product.Create(ctx, uid, product)
	if err != nil {
		return 0, checkAppError(s.log, err, "Create.Create")
	}

	return id, nil
}

func (s *ProductService) Update(ctx context.Context, uid int, product domain.Product) (err error) {
	if err = s.repo.Product.Update(ctx, uid, product); err != nil {
		return checkAppError(s.log, err, "Update.Update")
	}

	return nil
}

func (s *ProductService) Get(ctx context.Context, uid, id int) (domain.Product, error) {
	p, err := s.repo.Product.Get(ctx, uid, id)
	if err != nil {
		return domain.Product{}, checkAppError(s.log, err, "Update.Update")
	}

	return p, nil
}

func (s *ProductService) GetManyByFilter(ctx context.Context, filter domain.GetManyProductsFilter) (domain.GetManyProductsResponse, error) {
	res, err := s.repo.Product.GetManyByFilter(ctx, filter)
	if err != nil {
		return domain.GetManyProductsResponse{}, checkAppError(s.log, err, "GetManyByFilter.GetManyByFilter")
	}

	return res, nil
}
