package grpc

import (
	"context"
	"github.com/igntnk/stocky-2pc-controller/protobufs/sms_pb"
	"github.com/igntnk/stocky-sms/models"
	"github.com/igntnk/stocky-sms/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type supplyServer struct {
	sms_pb.UnimplementedSupplyServiceServer
	Logger        zerolog.Logger
	SupplyService service.SupplyService
}

func RegisterSupplyServer(grpcServer *grpc.Server, logger zerolog.Logger, supplyService service.SupplyService) {
	sms_pb.RegisterSupplyServiceServer(grpcServer, &supplyServer{
		Logger:        logger,
		SupplyService: supplyService,
	})
}

func (s *supplyServer) CreateSupply(ctx context.Context, req *sms_pb.CreateSupplyRequest) (*sms_pb.UuidResponse, error) {
	reqProducts := req.GetProducts()
	resProducts := make([]models.SupplyProduct, len(reqProducts))

	for i, product := range reqProducts {
		resProducts[i] = models.SupplyProduct{
			Product: models.Product{
				Uuid: product.ProductUuid,
			},
			Amount: float64(product.GetAmount()),
		}
	}

	resUuid, err := s.SupplyService.Create(ctx, models.SupplyWithProducts{
		Supply: models.Supply{
			Comment:         req.GetComment(),
			CreationDate:    time.Now().String(),
			DesiredDate:     req.GetDesiredDate(),
			Status:          models.Created,
			ResponsibleUser: req.GetResponsibleUser(),
			Edited:          false,
			EditedDate:      time.Now().String(),
			Cost:            float64(req.SupplyCost),
		},
		Products: resProducts,
	})
	if err != nil {
		return nil, err
	}

	return &sms_pb.UuidResponse{Uuid: resUuid}, nil
}

func (s *supplyServer) DeleteSupply(ctx context.Context, req *sms_pb.UuidRequest) (*sms_pb.UuidResponse, error) {
	res, err := s.SupplyService.Delete(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	return &sms_pb.UuidResponse{Uuid: res}, nil
}

func (s *supplyServer) UpdateSupplyInfo(ctx context.Context, req *sms_pb.UpdateSupplyInfoRequest) (*sms_pb.UuidResponse, error) {
	res, err := s.SupplyService.UpdateSupplyInfo(ctx, models.Supply{
		Uuid:            req.Uuid,
		Comment:         req.GetComment(),
		DesiredDate:     req.GetDesiredDate(),
		Status:          models.SupplyState(req.Status),
		ResponsibleUser: req.ResponsibleUser,
		Edited:          true,
		EditedDate:      time.Now().String(),
		Cost:            float64(req.Cost),
	})
	if err != nil {
		return nil, err
	}

	return &sms_pb.UuidResponse{Uuid: res}, nil
}

func (s *supplyServer) GetActiveSupplies(ctx context.Context, req *emptypb.Empty) (*sms_pb.GetActiveSuppliesResponse, error) {
	res, err := s.SupplyService.GetActiveSupplies(ctx)
	if err != nil {
		return nil, err
	}

	resModels := make([]*sms_pb.SupplyModel, len(res))
	for i, item := range res {
		resModels[i] = &sms_pb.SupplyModel{
			Uuid:            item.Uuid,
			Comment:         item.Comment,
			DesiredDate:     item.DesiredDate,
			Status:          string(item.Status),
			ResponsibleUser: item.ResponsibleUser,
			Cost:            float32(item.Cost),
		}
	}

	return &sms_pb.GetActiveSuppliesResponse{
		Supplies: resModels,
	}, nil
}

func (s *supplyServer) GetSupplyById(ctx context.Context, req *sms_pb.UuidRequest) (*sms_pb.SupplyModel, error) {
	res, err := s.SupplyService.GetSupplyById(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	return &sms_pb.SupplyModel{
		Uuid:            res.Uuid,
		Comment:         res.Comment,
		DesiredDate:     res.DesiredDate,
		Status:          string(res.Status),
		ResponsibleUser: res.ResponsibleUser,
		Cost:            float32(res.Cost),
	}, nil
}
