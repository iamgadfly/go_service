package project

import (
	"context"
	"go_service/config"
	"go_service/internal/models"
	"go_service/internal/project"
	"go_service/internal/project/repository"
	proto "go_service/pkg/api/project/proto"
	"time"
)

// ProjectServer ..
type ProjectServer struct {
	proto.UnimplementedProjectServer
	cfg     *config.Config
	useCase project.UseCase
	repo    *repository.ProjectRepo
}

// NewProjectServer ..
func NewProjectServer(cfg *config.Config, repo *repository.ProjectRepo, uc project.UseCase) *ProjectServer {
	return &ProjectServer{
		useCase: uc,
		cfg:     cfg,
		repo:    repo,
	}
}

// Create ..
func (p ProjectServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.ProjectResponse, error) {
	var projectRaw models.Project
	projectRaw.Name = req.GetName()
	projectCreate, err := p.useCase.Create(ctx, &projectRaw)
	if err != nil {
		return nil, err
	}

	return &proto.ProjectResponse{
		Id:        uint64(projectCreate.ID),
		Name:      projectCreate.Name,
		CreatedAt: projectCreate.CreatedAt.Format(time.DateTime),
	}, nil
}
