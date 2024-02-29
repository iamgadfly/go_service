package good

import (
	"context"
	"go_service/config"
	"go_service/internal/good"
	"go_service/internal/good/repository"
	"go_service/internal/models"
	proto "go_service/pkg/api/good/proto"
	"time"
)

// GoodServer ..
type GoodServer struct {
	proto.UnimplementedGoodServer
	cfg     *config.Config
	useCase good.UseCase
	repo    *repository.GoodRepo
}

// NewGoodServer ..
func NewGoodServer(cfg *config.Config, repo *repository.GoodRepo, uc good.UseCase) *GoodServer {
	return &GoodServer{
		useCase: uc,
		cfg:     cfg,
		repo:    repo,
	}
}

// Create ..
func (g GoodServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.GoodResponse, error) {
	goodCreate, err := g.useCase.Create(ctx, &models.Good{
		ProjectId: int(request.GetProjectId()),
		Name:      request.GetName(),
	})

	if err != nil {
		return nil, err
	}

	return createResp(goodCreate), nil
}

// Update ..
func (g GoodServer) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.GoodResponse, error) {
	goodUpdate, err := g.useCase.Update(ctx, &models.Good{
		ID:          int(req.GetId()),
		ProjectId:   int(req.ProjectId),
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return nil, err
	}

	return createResp(goodUpdate), nil
}

// Remove ..
func (g GoodServer) Remove(ctx context.Context, req *proto.RemoveRequest) (*proto.RemoveResponse, error) {
	goodRemove, err := g.useCase.Remove(ctx, &models.Good{
		ID:        int(req.GetId()),
		ProjectId: int(req.ProjectId),
	})

	if err != nil {
		return nil, err
	}

	return &proto.RemoveResponse{
		Id:         uint64(goodRemove.ID),
		CampaginId: uint64(goodRemove.CampaginId),
		Removed:    goodRemove.Removed,
	}, nil
}

// List ..
func (g GoodServer) List(ctx context.Context, req *proto.ListRequest) (*proto.ListResponse, error) {
	list, err := g.useCase.List(ctx, &models.GoodList{
		Limit:  int(req.GetLimit()),
		Offset: int(req.GetOffset()),
	})

	if err != nil {
		return nil, err
	}

	goods := make([]*proto.GoodResponse, len(list.Goods))
	for i, v := range list.Goods {
		goods[i] = createResp(&v)
	}

	return &proto.ListResponse{Meta: &proto.Meta{
		Total:   uint64(list.Total),
		Removed: uint64(list.Removed),
		Limit:   uint64(list.Limit),
		Offset:  uint64(list.Offset),
	}, Goods: goods}, nil
}

func createResp(good *models.Good) *proto.GoodResponse {
	return &proto.GoodResponse{
		Id:          uint64(good.ID),
		ProjectId:   uint64(good.ProjectId),
		Name:        good.Name,
		Description: good.Description,
		Priority:    int64(good.Priority),
		Removed:     good.Removed,
		CreatedAt:   good.CreatedAt.Format(time.DateTime),
	}
}
