package http

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go_service/config"
	"go_service/internal/models"
	"go_service/internal/project"
	"go_service/pkg/httpErrors"
	"go_service/pkg/utils"
	"net/http"
)

type projectHandler struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
	uc     project.UseCase
}

// NewProjectHandler ..
func NewProjectHandler(cfg *config.Config, logger *zap.SugaredLogger, uc project.UseCase) projectHandler {
	return projectHandler{
		cfg:    cfg,
		logger: logger,
		uc:     uc,
	}
}

// Create
// @Summary Create project
// @tags projects
// @description Create project
// @Param data body models.CreateProjectReq true "data for create"
// @success 200 {object} models.Project
// @router /api/v1/project/create [post]
func (h projectHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "projectHandler.Create")
		defer span.Finish()
		p := models.Project{}
		if err := c.Bind(&p); err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		projectCreated, er := h.uc.Create(ctx, &p)
		if er != nil {
			return c.JSON(httpErrors.ErrorResponse(er))
		}

		return c.JSON(http.StatusOK, projectCreated)
	}
}
