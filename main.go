package main

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/pgxv5"
	"github.com/igntnk/stocky-sms/config"
	grpcapp "github.com/igntnk/stocky-sms/grpc"
	"github.com/igntnk/stocky-sms/repository"
	"github.com/igntnk/stocky-sms/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
	"os/signal"
	"syscall"

	"context"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	mainCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := config.Get(logger)

	dbConf, err := pgxpool.ParseConfig(cfg.Database.URI)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse database config")
		return
	}

	pool, err := pgxpool.NewWithConfig(mainCtx, dbConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
		return
	}

	db := stdlib.OpenDBFromPool(pool)

	err = goose.SetDialect("postgres")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to set postgres dialect")
		return
	}

	err = goose.Up(db, "cmd/changelog")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to migrate database")
		return
	}

	productRepo := repository.NewProductRepository(logger, pool, trmpgx.DefaultCtxGetter)
	supplyRepo := repository.NewSupplyRepository(logger, pool, trmpgx.DefaultCtxGetter)

	productService := service.NewProductService(logger, productRepo)
	supplyService := service.NewSupplyService(logger, supplyRepo)

	grpcServer := grpc.NewServer()
	grpcapp.RegisterSupplyServer(grpcServer, logger, supplyService)
	grpcapp.RegisterProductServer(grpcServer, logger, productService)

	cookedGrpcServer := grpcapp.New(grpcServer, cfg.Server.Port, logger)
	go func() {
		cookedGrpcServer.MustRun()
	}()

	select {
	case <-mainCtx.Done():
		logger.Info().Msg("shutting down")
		cookedGrpcServer.Stop()
	}
}
