package usecase

import (
	"context"
	"go.uber.org/zap"
	"go_service/config"
	"go_service/internal/models"
	"go_service/internal/project"
)

type ProjectUC struct {
	cfg         *config.Config
	logger      *zap.SugaredLogger
	projectRepo project.Repository
}

// NewProjectUC ..
func NewProjectUC(cfg *config.Config, logger *zap.SugaredLogger, projectRepo project.Repository) project.UseCase {
	return &ProjectUC{
		cfg:         cfg,
		logger:      logger,
		projectRepo: projectRepo,
	}
}

// Create ..
func (u *ProjectUC) Create(ctx context.Context, project *models.Project) (*models.Project, error) {
	g, err := u.projectRepo.Create(ctx, project)
	if err != nil {
		return nil, err
	}

	return g, nil
}
