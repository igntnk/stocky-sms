package service

import (
	"context"
	"github.com/igntnk/stocky_sms/models"
	"github.com/igntnk/stocky_sms/repository"
	"github.com/rs/zerolog"
)

type ProductService interface {
	Create(ctx context.Context, productCode string) (string, error)
	Delete(ctx context.Context, uuid string) (string, error)
	SetStoreCost(ctx context.Context, product models.Product) error
	SetStoreAmount(ctx context.Context, product models.Product) error
	GetStoreAmount(ctx context.Context, uuid string) (float64, error)
}

func NewProductService(logger zerolog.Logger, repo repository.ProductRepository) ProductService {
	return &productService{
		logger: logger,
		repo:   repo,
	}
}

type productService struct {
	logger zerolog.Logger
	repo   repository.ProductRepository
}

func (p productService) Create(ctx context.Context, productCode string) (string, error) {
	return p.repo.Create(ctx, productCode)
}

func (p productService) Delete(ctx context.Context, uuid string) (string, error) {
	return p.repo.Delete(ctx, uuid)
}

func (p productService) SetStoreCost(ctx context.Context, product models.Product) error {
	return p.repo.SetStoreCost(ctx, product)
}

func (p productService) SetStoreAmount(ctx context.Context, product models.Product) error {
	return p.repo.SetStoreAmount(ctx, product)
}

func (p productService) GetStoreAmount(ctx context.Context, uuid string) (float64, error) {
	return p.repo.GetStoreAmount(ctx, uuid)
}
