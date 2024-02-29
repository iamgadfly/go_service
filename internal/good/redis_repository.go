package good

import (
	"context"
	"go_service/internal/models"
)

type RedisRepository interface {
	GetGoodsByIDCtx(ctx context.Context, key string) (*models.Good, error)
	SetGoodsCtx(ctx context.Context, key string, seconds int, goods *models.Good) error
	DeleteGoodsCtx(ctx context.Context, key string) error
}
