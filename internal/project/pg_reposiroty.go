package project

import (
	"context"
	"go_service/internal/models"
)

type Repository interface {
	Create(ctx context.Context, good *models.Project) (*models.Project, error)
}
