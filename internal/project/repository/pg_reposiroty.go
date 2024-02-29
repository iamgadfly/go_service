package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go_service/internal/models"
)

type ProjectRepo struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

// NewProjectRepo ..
func NewProjectRepo(db *sqlx.DB, logger *zap.SugaredLogger) *ProjectRepo {
	return &ProjectRepo{
		db:     db,
		logger: logger,
	}
}

func (r ProjectRepo) Create(ctx context.Context, project *models.Project) (*models.Project, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "projectRepo.Create")
	defer span.Finish()

	p := models.Project{}
	if err := r.db.QueryRowxContext(
		ctx,
		Create,
		&project.Name,
	).StructScan(&p); err != nil {
		return nil, errors.Wrap(err, "projectRepo.Create.QueryRowxContext")
	}
	return &p, nil
}
