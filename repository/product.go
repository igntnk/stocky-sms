package repository

import (
	"context"
	"github.com/igntnk/stocky_sms/models"
)

type ProductRepository interface {
	Create(ctx context.Context, product models.Product) (string, error)
	Delete(ctx context.Context, product models.Product) (string, error)
	SetStoreCost(ctx context.Context, cost float64) error
	SetStoreAmount(ctx context.Context, amount float64) error
	GetStoreAmount(ctx context.Context, uuid string) (float64, error)
}
