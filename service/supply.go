package service

import (
	"context"
	"github.com/igntnk/stocky-sms/models"
	"github.com/igntnk/stocky-sms/repository"
	"github.com/rs/zerolog"
)

type SupplyService interface {
	Create(ctx context.Context, supply models.SupplyWithProducts) (string, error)
	Delete(ctx context.Context, uuid string) (string, error)
	UpdateSupplyInfo(ctx context.Context, supply models.Supply) (string, error)
	GetActiveSupplies(ctx context.Context) ([]models.SupplyWithProducts, error)
	GetSupplyById(ctx context.Context, uuid string) (models.SupplyWithProducts, error)
}

func NewSupplyService(logger zerolog.Logger, repo repository.SupplyRepository) SupplyService {
	return &supplyService{
		logger: logger,
		repo:   repo,
	}
}

type supplyService struct {
	logger zerolog.Logger
	repo   repository.SupplyRepository
}

func (s supplyService) Create(ctx context.Context, supply models.SupplyWithProducts) (string, error) {
	return s.repo.Create(ctx, supply)
}

func (s supplyService) Delete(ctx context.Context, uuid string) (string, error) {
	return s.repo.Delete(ctx, uuid)
}

func (s supplyService) UpdateSupplyInfo(ctx context.Context, supply models.Supply) (string, error) {
	return s.repo.UpdateSupplyInfo(ctx, supply)
}

func (s supplyService) GetActiveSupplies(ctx context.Context) ([]models.SupplyWithProducts, error) {
	return s.repo.GetActiveSupplies(ctx)
}

func (s supplyService) GetSupplyById(ctx context.Context, uuid string) (models.SupplyWithProducts, error) {
	return s.repo.GetSupplyById(ctx, uuid)
}
