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
		// e :=
		return 0, checkAppError(
			"Create.Create",
			err,
			s.log.Error().
				Int("uid", uid),
		)
	}

	return id, nil
}

func (s *ProductService) Update(ctx context.Context, uid int, product domain.Product) (err error) {
	if err = s.repo.Product.Update(ctx, uid, product); err != nil {
		return checkAppError(
			"Update.Update",
			err,
			s.log.Error().
				Int("uid", uid).
				Int("product id", product.ID),
		)
	}

	return nil
}

func (s *ProductService) Get(ctx context.Context, uid, id int) (domain.Product, error) {
	p, err := s.repo.Product.Get(ctx, uid, id)
	if err != nil {
		return domain.Product{}, checkAppError(
			"Get.Get",
			err,
			s.log.Error().
				Int("uid", uid).
				Int("product id", id),
		)
	}

	return p, nil
}

func (s *ProductService) GetManyByFilter(ctx context.Context, filter domain.GetManyProductsFilter) (domain.GetManyProductsResponse, error) {
	res, err := s.repo.Product.GetManyByFilter(ctx, filter)
	if err != nil {
		return domain.GetManyProductsResponse{}, checkAppError(
			"GetManyByFilter.GetManyByFilter",
			err,
			s.log.Error().
				Interface("filter", filter),
		)
	}

	return res, nil
}
