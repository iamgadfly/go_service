package repository

import (
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go_service/internal/good"
	"go_service/internal/models"
	"time"
)

type GoodRedisRepository struct {
	logger      *zap.SugaredLogger
	redisClient *redis.Client
}

// NewGoodRedisRepository ..
func NewGoodRedisRepository(logger *zap.SugaredLogger, r *redis.Client) good.RedisRepository {
	return GoodRedisRepository{
		logger:      logger,
		redisClient: r,
	}
}

func (r GoodRedisRepository) GetGoodsByIDCtx(ctx context.Context, key string) (*models.Good, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "newsRedisRepo.GetNewsByIDCtx")
	defer span.Finish()

	goodsBytes, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "newsRedisRepo.GetNewsByIDCtx.redisClient.Get")
	}

	g := &models.Good{}
	if err = json.Unmarshal(goodsBytes, g); err != nil {
		return nil, errors.Wrap(err, "newsRedisRepo.GetNewsByIDCtx.json.Unmarshal")
	}

	return g, nil
}

// SetGoodsCtx ..
func (r GoodRedisRepository) SetGoodsCtx(ctx context.Context, key string, seconds int, goods *models.Good) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "goodRedisRepository.SetGoodsCtx")
	defer span.Finish()

	goodsBytes, err := json.Marshal(goods)
	if err != nil {
		return errors.Wrap(err, "goodRedisRepository.SetGoodsCtx.json.Marshal")
	}
	if err = r.redisClient.Set(ctx, key, goodsBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "goodRedisRepository.SetGoodsCtx.redisClient.Set")
	}
	return nil
}

// DeleteGoodsCtx ..
func (r GoodRedisRepository) DeleteGoodsCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "goodRedisRepository.DeleteGoodsCtx")
	defer span.Finish()

	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "goodRedisRepository.DeleteGoodsCtx.redisClient.Del")
	}
	return nil
}
