package service

import (
	"context"
	"github.com/igntnk/stocky-sms/models"
	"github.com/igntnk/stocky-sms/repository"
	"github.com/rs/zerolog"
)

type ProductService interface {
	Create(ctx context.Context, storeCost float64) (string, error)
	Delete(ctx context.Context, uuid string) (string, error)
	SetStoreCost(ctx context.Context, product models.Product) error
	SetStoreAmount(ctx context.Context, product models.Product) error
	GetStoreAmount(ctx context.Context, uuid string) (float64, error)
	RemoveCoupleProducts(ctx context.Context, products []models.RemoveProductRequest) ([]string, error)
	WriteOnCoupleProducts(ctx context.Context, products []models.RemoveProductRequest) ([]string, error)
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

func (p productService) RemoveCoupleProducts(ctx context.Context, products []models.RemoveProductRequest) ([]string, error) {
	return p.repo.RemoveCoupleProducts(ctx, products)
}

func (p productService) WriteOnCoupleProducts(ctx context.Context, products []models.RemoveProductRequest) ([]string, error) {
	return p.repo.WriteOnCoupleProducts(ctx, products)
}

func (p productService) Create(ctx context.Context, storeCost float64) (string, error) {
	return p.repo.Create(ctx, storeCost)
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
