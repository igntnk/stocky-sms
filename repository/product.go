package repository

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/google/uuid"
	"github.com/igntnk/stocky-sms/db"
	"github.com/igntnk/stocky-sms/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type ProductRepository interface {
	Create(ctx context.Context, productCode string) (string, error)
	Delete(ctx context.Context, uuid string) (string, error)
	SetStoreCost(ctx context.Context, product models.Product) error
	SetStoreAmount(ctx context.Context, product models.Product) error
	GetStoreAmount(ctx context.Context, uuid string) (float64, error)
}

func NewProductRepository(logger zerolog.Logger, pool *pgxpool.Pool, getter *trmpgx.CtxGetter) ProductRepository {
	return &productRepository{
		baseRepository{
			Logger: logger,
			DB:     pool,
			Getter: getter,
		},
	}
}

type productRepository struct {
	baseRepository
}

func (p productRepository) Create(ctx context.Context, productCode string) (string, error) {
	conn := p.Getter.DefaultTrOrDB(ctx, p.DB)
	q := db.New(conn)

	res, err := q.CreateProduct(ctx, productCode)
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func (p productRepository) Delete(ctx context.Context, productUuid string) (string, error) {
	conn := p.Getter.DefaultTrOrDB(ctx, p.DB)
	q := db.New(conn)

	prUuid, err := uuid.Parse(productUuid)
	if err != nil {
		return "", err
	}

	res, err := q.DeleteProduct(ctx, pgtype.UUID{
		Bytes: prUuid,
		Valid: true,
	})

	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func (p productRepository) SetStoreCost(ctx context.Context, product models.Product) error {
	conn := p.Getter.DefaultTrOrDB(ctx, p.DB)
	q := db.New(conn)

	prUuid, err := uuid.Parse(product.Uuid)
	if err != nil {
		return err
	}

	num, err := Float64ToNumericWithPrecision(product.StoreCost, 64)
	if err != nil {
		return err
	}

	return q.SetStoreCost(ctx, db.SetStoreCostParams{
		StoreCost: num,
		Uuid: pgtype.UUID{
			Bytes: prUuid,
			Valid: true,
		},
	})
}

func (p productRepository) SetStoreAmount(ctx context.Context, product models.Product) error {
	conn := p.Getter.DefaultTrOrDB(ctx, p.DB)
	q := db.New(conn)

	prUuid, err := uuid.Parse(product.Uuid)
	if err != nil {
		return err
	}

	num, err := Float64ToNumericWithPrecision(product.StoreAmount, 64)
	if err != nil {
		return err
	}

	return q.SetStoreAmount(ctx, db.SetStoreAmountParams{
		StoreAmount: num,
		Uuid: pgtype.UUID{
			Bytes: prUuid,
			Valid: true,
		},
	})
}

func (p productRepository) GetStoreAmount(ctx context.Context, productUuid string) (float64, error) {
	conn := p.Getter.DefaultTrOrDB(ctx, p.DB)
	q := db.New(conn)

	prUuid, err := uuid.Parse(productUuid)
	if err != nil {
		return 0, err
	}

	res, err := q.GetStoreAmount(ctx, pgtype.UUID{
		Bytes: prUuid,
		Valid: true,
	})
	if err != nil {
		return 0, err
	}

	return NumericToFloat64(res)
}
