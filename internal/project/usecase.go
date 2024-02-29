package project

import (
	"context"
	"go_service/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, good *models.Project) (*models.Project, error)
}
