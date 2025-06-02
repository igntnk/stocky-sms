package repository

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/google/uuid"
	"github.com/igntnk/stocky-sms/db"
	"github.com/igntnk/stocky-sms/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

type SupplyRepository interface {
	Create(ctx context.Context, supply models.SupplyWithProducts) (string, error)
	Delete(ctx context.Context, uuid string) (string, error)
	UpdateSupplyInfo(ctx context.Context, supply models.Supply) (string, error)
	GetActiveSupplies(ctx context.Context) ([]models.SupplyWithProducts, error)
	GetSupplyById(ctx context.Context, uuid string) (models.SupplyWithProducts, error)
}

func NewSupplyRepository(logger zerolog.Logger, pool *pgxpool.Pool, getter *trmpgx.CtxGetter) SupplyRepository {
	return &supplyRepository{
		baseRepository{
			Logger: logger,
			DB:     pool,
			Getter: getter,
		},
	}
}

type supplyRepository struct {
	baseRepository
}

func (s supplyRepository) Create(ctx context.Context, supply models.SupplyWithProducts) (string, error) {
	conn := s.Getter.DefaultTrOrDB(ctx, s.DB)
	q := db.New(conn)

	cost, err := Float64ToNumericWithPrecision(supply.Cost, 64)
	if err != nil {
		return "", err
	}

	desTime, err := time.Parse(time.RFC3339, supply.DesiredDate)
	if err != nil {
		return "", err
	}

	supUuid, err := q.CreateSupply(ctx, db.CreateSupplyParams{
		Comment: pgtype.Text{
			String: supply.Comment,
			Valid:  true,
		},
		DesiredDate: pgtype.Timestamp{
			Time:             desTime,
			InfinityModifier: 0,
			Valid:            true,
		},
		ResponsibleUser: supply.ResponsibleUser,
		EditedDate: pgtype.Timestamp{
			Time:             time.Now(),
			InfinityModifier: 0,
			Valid:            true,
		},
		Cost: cost,
	})

	batch := pgx.Batch{}

	for _, product := range supply.Products {
		batch.Queue("insert into supply_product (product_uuid, supply_uuid, amount) values ($1, $2, $3)", product.Uuid, supUuid.String(), product.Amount)
	}

	results := s.DB.SendBatch(ctx, &batch)
	defer results.Close()

	for range supply.Products {
		_, err := results.Exec()
		if err != nil {
			return "", err
		}
	}

	return supUuid.String(), nil
}

func (s supplyRepository) Delete(ctx context.Context, supUuid string) (string, error) {
	conn := s.Getter.DefaultTrOrDB(ctx, s.DB)
	q := db.New(conn)

	sUuid, err := uuid.Parse(supUuid)
	if err != nil {
		return "", err
	}

	res, err := q.DeleteSupply(ctx, pgtype.UUID{
		Bytes: sUuid,
		Valid: true,
	})
	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func (s supplyRepository) UpdateSupplyInfo(ctx context.Context, supply models.Supply) (string, error) {
	conn := s.Getter.DefaultTrOrDB(ctx, s.DB)
	q := db.New(conn)

	desTime, err := time.Parse(time.RFC3339, supply.DesiredDate)
	if err != nil {
		return "", err
	}

	sUuid, err := uuid.Parse(supply.Uuid)
	if err != nil {
		return "", err
	}

	num, err := Float64ToNumericWithPrecision(supply.Cost, 64)
	if err != nil {
		return "", err
	}

	err = q.UpdateSupplyInfo(ctx, db.UpdateSupplyInfoParams{
		Comment: pgtype.Text{
			String: supply.Comment,
			Valid:  true,
		},
		DesiredDate: pgtype.Timestamp{
			Time:             desTime,
			InfinityModifier: 0,
			Valid:            true,
		},
		Status:          db.SupplyStatus(supply.Status),
		ResponsibleUser: supply.ResponsibleUser,
		Cost:            num,
		Uuid: pgtype.UUID{
			Bytes: sUuid,
			Valid: true,
		},
	})
	if err != nil {
		return "", err
	}

	return supply.Uuid, nil
}

func (s supplyRepository) GetActiveSupplies(ctx context.Context) ([]models.SupplyWithProducts, error) {
	conn := s.Getter.DefaultTrOrDB(ctx, s.DB)
	q := db.New(conn)

	supplies, err := q.GetActiveSupplies(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]models.SupplyWithProducts, len(supplies))
	for i, supply := range supplies {
		num, err := NumericToFloat64(supply.Cost)
		if err != nil {
			return nil, err
		}

		res[i] = models.SupplyWithProducts{
			Supply: models.Supply{
				Uuid:            supply.Uuid.String(),
				Comment:         supply.Comment.String,
				CreationDate:    supply.CreationDate.Time.String(),
				DesiredDate:     supply.DesiredDate.Time.String(),
				Status:          models.SupplyState(supply.Status),
				ResponsibleUser: supply.ResponsibleUser,
				Edited:          supply.Edited.Bool,
				EditedDate:      supply.EditedDate.Time.String(),
				Cost:            num,
			},
		}
	}

	return res, nil
}

func (s supplyRepository) GetSupplyById(ctx context.Context, supUuid string) (models.SupplyWithProducts, error) {
	conn := s.Getter.DefaultTrOrDB(ctx, s.DB)
	q := db.New(conn)

	sUuid, err := uuid.Parse(supUuid)
	if err != nil {
		return models.SupplyWithProducts{}, err
	}

	supply, err := q.GetSupplyById(ctx, pgtype.UUID{
		Bytes: sUuid,
		Valid: true,
	})
	if err != nil {
		return models.SupplyWithProducts{}, err
	}

	num, err := NumericToFloat64(supply.Cost)
	if err != nil {
		return models.SupplyWithProducts{}, err
	}

	return models.SupplyWithProducts{
		Supply: models.Supply{
			Uuid:            supply.Uuid.String(),
			Comment:         supply.Comment.String,
			CreationDate:    supply.CreationDate.Time.String(),
			DesiredDate:     supply.DesiredDate.Time.String(),
			Status:          models.SupplyState(supply.Status),
			ResponsibleUser: supply.ResponsibleUser,
			Edited:          supply.Edited.Bool,
			EditedDate:      supply.EditedDate.Time.String(),
			Cost:            num,
		},
	}, nil
}
