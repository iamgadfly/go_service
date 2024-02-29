package usecase

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go_service/config"
	"go_service/internal/good"
	"go_service/internal/models"
	"strconv"
)

type GoodUC struct {
	cfg       *config.Config
	logger    *zap.SugaredLogger
	redisRepo good.RedisRepository
	goodsRepo good.Repository
}

const (
	basePrefix    = "api-goods:"
	cacheDuration = 60
)

// NewGoodUC ..
func NewGoodUC(cfg *config.Config, logger *zap.SugaredLogger, redisRepo good.RedisRepository, goodsRepo good.Repository) good.UseCase {
	return &GoodUC{
		cfg:       cfg,
		logger:    logger,
		redisRepo: redisRepo,
		goodsRepo: goodsRepo,
	}
}

// Create ..
func (u *GoodUC) Create(ctx context.Context, good *models.Good) (*models.Good, error) {
	g, err := u.goodsRepo.Create(ctx, good)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.SetGoodsCtx(ctx, u.getKeyWithPrefix(strconv.Itoa(g.ID)), cacheDuration, g); err != nil {
		u.logger.Errorf("newsUC.GetNewsByID.SetNewsCtx: %s", err)
	}

	return g, nil
}

// Update ..
func (u *GoodUC) Update(ctx context.Context, good *models.Good) (*models.Good, error) {
	g, err := u.goodsRepo.Update(ctx, good)

	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.DeleteGoodsCtx(ctx, u.getKeyWithPrefix(string(rune(g.ID)))); err != nil {
		u.logger.Errorf("goodUC.Update.DeleteGoodsCtx: %v", err)
	}

	return g, nil
}

// FindById ..
func (u *GoodUC) FindById(ctx context.Context, good *models.Good) (*models.Good, error) {
	g, err := u.goodsRepo.FindById(ctx, good.ID)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// Remove ..
func (u *GoodUC) Remove(ctx context.Context, good *models.Good) (*models.RemoveResp, error) {
	g, err := u.goodsRepo.Remove(ctx, good)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.DeleteGoodsCtx(ctx, u.getKeyWithPrefix(string(rune(g.ID)))); err != nil {
		u.logger.Errorf("goodUC.Update.DeleteGoodsCtx: %v", err)
	}

	return g, nil
}

// List ..
func (u *GoodUC) List(ctx context.Context, list *models.GoodList) (*models.GoodList, error) {
	l, err := u.goodsRepo.List(ctx, list)
	if err != nil {
		return nil, err
	}

	return l, nil
}

// Reprioritiize ..
func (u *GoodUC) Reprioritiize(ctx context.Context, good *models.Good) (*[]models.PriorityResp, error) {
	g, err := u.goodsRepo.FindById(ctx, good.ID)
	if err != nil {
		return nil, err
	}
	g.Priority = good.Priority

	res, err := u.goodsRepo.Reprioritiize(ctx, g)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// getKeyWithPrefix ..
func (u *GoodUC) getKeyWithPrefix(goodsID string) string {
	return fmt.Sprintf("%s:%s", basePrefix, goodsID)
}
