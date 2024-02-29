package good

import (
	"context"
	"go_service/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, good *models.Good) (*models.Good, error)
	Update(ctx context.Context, good *models.Good) (*models.Good, error)
	Remove(ctx context.Context, good *models.Good) (*models.RemoveResp, error)
	List(ctx context.Context, list *models.GoodList) (*models.GoodList, error)
	Reprioritiize(ctx context.Context, good *models.Good) (*[]models.PriorityResp, error)
}
