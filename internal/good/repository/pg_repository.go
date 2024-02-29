package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go_service/internal/models"
)

type GoodRepo struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

// NewGoodRepo ..
func NewGoodRepo(logger *zap.SugaredLogger, db *sqlx.DB) *GoodRepo {
	return &GoodRepo{
		logger: logger,
		db:     db,
	}
}

// Create ..
func (r GoodRepo) Create(ctx context.Context, good *models.Good) (*models.Good, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "goodRepo.Create")
	defer span.Finish()

	g := models.Good{}
	if err := r.db.QueryRowxContext(
		ctx,
		Create,
		&good.ProjectId,
		&good.Name,
		&good.Description,
		&good.Priority,
		&good.Removed,
	).StructScan(&g); err != nil {
		return nil, errors.Wrap(err, "goodRepo.Create.QueryRowxContext")
	}

	return &g, nil
}

// Update ..
func (r GoodRepo) Update(ctx context.Context, good *models.Good) (*models.Good, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "goodRepo.Update")
	defer span.Finish()

	goodCheck, err := r.FindById(ctx, good.ID)
	if err != nil {
		r.logger.Info(err)
		return nil, errors.Wrap(err, "goodRepo.Update.FindById")
	}
	if goodCheck.ProjectId != good.ProjectId {
		return nil, errors.New("project_id is not fits")
	}

	g := models.Good{}
	if err := r.db.QueryRowxContext(
		ctx,
		Update,
		&good.ID,
		&good.Name,
		&good.Description,
	).StructScan(&g); err != nil {
		return nil, errors.Wrap(err, "goodRepo.Update.QueryRowxContext")
	}

	return &g, nil
}

// FindById ..
func (r GoodRepo) FindById(ctx context.Context, goodId int) (*models.Good, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "goodRepo.GetNewsByID")
	defer span.Finish()

	g := &models.Good{}
	if err := r.db.GetContext(ctx, g, FindById, &goodId); err != nil {
		r.logger.Info(err)
		return nil, errors.Wrap(err, "goodRepo.FindById.GetContext")
	}

	return g, nil
}

// Remove ..
func (r GoodRepo) Remove(ctx context.Context, good *models.Good) (*models.RemoveResp, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "goodRepo.Remove")
	defer span.Finish()

	goodCheck, err := r.FindById(ctx, good.ID)
	if err != nil {
		r.logger.Info(err)
		return nil, errors.Wrap(err, "goodRepo.Remove.FindById")
	}
	if goodCheck.ProjectId != good.ProjectId {
		return nil, errors.New("project_id is not fits")
	}

	g := models.Good{}
	if err := r.db.QueryRowxContext(
		ctx,
		Remove,
		&good.ID,
	).StructScan(&g); err != nil {
		return nil, errors.Wrap(err, "goodRepo.Remove.QueryRowxContext")
	}

	return &models.RemoveResp{
		ID:         g.ID,
		CampaginId: g.ProjectId,
		Removed:    g.Removed,
	}, nil
}

// List ..
func (r GoodRepo) List(ctx context.Context, list *models.GoodList) (*models.GoodList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "goodRepo.Remove")
	defer span.Finish()

	var goods []models.Good
	if list.Limit == 0 {
		err := r.db.SelectContext(ctx, &goods, ListWithoutLimit, list.Offset)
		if err != nil {
			r.logger.Info(err)
			return nil, err
		}
	} else {
		err := r.db.SelectContext(ctx, &goods, List, list.Offset, list.Limit)
		if err != nil {
			r.logger.Info(err)
			return nil, err
		}
	}
	list.Total = len(goods)
	list.Goods = goods
	list.Removed = goods[0].RemovedCount

	return list, nil

}

// Reprioritiize ..
func (r GoodRepo) Reprioritiize(ctx context.Context, good *models.Good) (*[]models.PriorityResp, error) {

	g := models.Good{}
	if err := r.db.QueryRowxContext(
		ctx,
		UpdatePriority,
		&good.Priority,
		&good.ID,
	).StructScan(&g); err != nil {
		return nil, errors.Wrap(err, "goodRepo.Reprioritiize.QueryRowxContext")
	}

	r.db.QueryRowxContext(ctx, Reprioritiize)

	var res []models.PriorityResp
	err := r.db.SelectContext(ctx, &res, GetReprioritiize, &good.ID)
	if err != nil {
		r.logger.Info(err)
		return nil, err
	}

	return &res, nil
}
