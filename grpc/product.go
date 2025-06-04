package grpc

import (
	"context"
	"github.com/igntnk/stocky-2pc-controller/protobufs/sms_pb"
	"github.com/igntnk/stocky-sms/models"
	"github.com/igntnk/stocky-sms/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"sync"
)

type productServer struct {
	sms_pb.UnimplementedProductServiceServer
	Logger           zerolog.Logger
	ProductService   service.ProductService
	changeProductsMu sync.Mutex
}

func RegisterProductServer(server *grpc.Server, logger zerolog.Logger, productService service.ProductService) {
	sms_pb.RegisterProductServiceServer(server, &productServer{Logger: logger, ProductService: productService, changeProductsMu: sync.Mutex{}})
}

func (s *productServer) ChangeCoupleProductAmount(stream grpc.BidiStreamingServer[sms_pb.RemoveProductsRequest, sms_pb.CoupleUuidResponse]) error {
	// lock resources event
	_, err := stream.Recv()
	if err != nil {
		return err
	}

	s.changeProductsMu.Lock()
	defer s.changeProductsMu.Unlock()

	ctx := stream.Context()

	changeProductRequest, err := stream.Recv()
	if err != nil {
		return err
	}

	res, err := s.RemoveCoupleProducts(ctx, changeProductRequest)
	if err != nil {
		return err
	}

	err = stream.Send(res)
	if err != nil {
		return err
	}

	return nil
}

func (s *productServer) CreateProduct(ctx context.Context, req *sms_pb.CreateProductMessage) (*sms_pb.UuidResponse, error) {
	res, err := s.ProductService.Create(ctx, float64(req.StoreCost))
	if err != nil {
		return nil, err
	}

	return &sms_pb.UuidResponse{Uuid: res}, nil
}

func (s *productServer) DeleteProduct(ctx context.Context, req *sms_pb.UuidRequest) (*sms_pb.UuidResponse, error) {
	res, err := s.ProductService.Delete(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &sms_pb.UuidResponse{Uuid: res}, nil
}

func (s *productServer) SetStoreCost(ctx context.Context, req *sms_pb.SetProductCostRequest) (*sms_pb.UuidResponse, error) {
	err := s.ProductService.SetStoreCost(ctx, models.Product{
		Uuid:      req.Uuid,
		StoreCost: float64(req.StoreCost),
	})
	if err != nil {
		return nil, err
	}

	return &sms_pb.UuidResponse{Uuid: req.Uuid}, nil
}

func (s *productServer) SetStoreAmount(ctx context.Context, req *sms_pb.SetProductAmountRequest) (*sms_pb.UuidResponse, error) {
	err := s.ProductService.SetStoreAmount(ctx, models.Product{
		Uuid:        req.Uuid,
		StoreAmount: float64(req.StoreAmount),
	})
	if err != nil {
		return nil, err
	}

	return &sms_pb.UuidResponse{Uuid: req.Uuid}, nil
}

func (s *productServer) GetStoreAmount(ctx context.Context, req *sms_pb.UuidRequest) (*sms_pb.GetStoreAmountResponse, error) {
	res, err := s.ProductService.GetStoreAmount(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &sms_pb.GetStoreAmountResponse{
		StoreAmount: float32(res),
	}, nil
}

func (s *productServer) RemoveCoupleProducts(ctx context.Context, req *sms_pb.RemoveProductsRequest) (*sms_pb.CoupleUuidResponse, error) {
	products := make([]models.RemoveProductRequest, len(req.Products))
	for i, product := range req.Products {
		products[i] = models.RemoveProductRequest{
			Uuid:   product.Uuid,
			Amount: float64(product.StoreAmount),
		}
	}

	res, err := s.ProductService.RemoveCoupleProducts(ctx, products)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(res))
	for i, product := range res {
		result[i] = product
	}

	return &sms_pb.CoupleUuidResponse{Uuids: result}, nil
}

func (s *productServer) WriteOnCoupleProducts(ctx context.Context, req *sms_pb.RemoveProductsRequest) (*sms_pb.CoupleUuidResponse, error) {
	products := make([]models.RemoveProductRequest, len(req.Products))
	for i, product := range req.Products {
		products[i] = models.RemoveProductRequest{
			Uuid:   product.Uuid,
			Amount: float64(product.StoreAmount),
		}
	}

	res, err := s.ProductService.WriteOnCoupleProducts(ctx, products)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(res))
	for i, product := range res {
		result[i] = product
	}

	return &sms_pb.CoupleUuidResponse{Uuids: result}, nil
}
