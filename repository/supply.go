package repository

import (
	"context"
	"github.com/igntnk/stocky_sms/models"
)

type Supply interface {
	Create(ctx context.Context, supply models.SupplyWithProducts) (string, error)
	Delete(ctx context.Context, uuid string) (string, error)
	UpdateSupplyInfo(ctx context.Context, supply models.Supply) (string, error)
	UpdateSupplyProducts(ctx context.Context, products []models.Product) (string, error)
	GetActiveSupplies(ctx context.Context) ([]models.SupplyWithProducts, error)
	GetSupplyById(ctx context.Context, uuid string) (models.SupplyWithProducts, error)
}
