package repository

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type baseRepository struct {
	Logger zerolog.Logger
	DB     *pgxpool.Pool
	Getter *trmpgx.CtxGetter
}
