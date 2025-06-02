package grpc

import (
	"context"
	"github.com/igntnk/stocky_sms/models"
	"github.com/igntnk/stocky_sms/proto/pb"
	"github.com/igntnk/stocky_sms/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type productServer struct {
	pb.UnimplementedProductServiceServer
	Logger         zerolog.Logger
	ProductService service.ProductService
}

func RegisterProductServer(server *grpc.Server, logger zerolog.Logger, productService service.ProductService) {
	pb.RegisterProductServiceServer(server, &productServer{Logger: logger, ProductService: productService})
}

func (s *productServer) CreateProduct(ctx context.Context, req *pb.CreateProductMessage) (*pb.UuidResponse, error) {
	res, err := s.ProductService.Create(ctx, req.ProductCode)
	if err != nil {
		return nil, err
	}

	return &pb.UuidResponse{Uuid: res}, nil
}

func (s *productServer) DeleteProduct(ctx context.Context, req *pb.UuidRequest) (*pb.UuidResponse, error) {
	res, err := s.ProductService.Delete(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.UuidResponse{Uuid: res}, nil
}

func (s *productServer) SetStoreCost(ctx context.Context, req *pb.SetProductCostRequest) (*pb.UuidResponse, error) {
	err := s.ProductService.SetStoreCost(ctx, models.Product{
		Uuid:      req.Uuid,
		StoreCost: float64(req.StoreCost),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UuidResponse{Uuid: req.Uuid}, nil
}

func (s *productServer) SetStoreAmount(ctx context.Context, req *pb.SetProductAmountRequest) (*pb.UuidResponse, error) {
	err := s.ProductService.SetStoreCost(ctx, models.Product{
		Uuid:        req.Uuid,
		StoreAmount: float64(req.StoreAmount),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UuidResponse{Uuid: req.Uuid}, nil
}

func (s *productServer) GetStoreAmount(ctx context.Context, req *pb.UuidRequest) (*pb.GetStoreAmountResponse, error) {
	res, err := s.ProductService.GetStoreAmount(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.GetStoreAmountResponse{
		StoreAmount: float32(res),
	}, nil
}
