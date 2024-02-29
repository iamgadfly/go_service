package http

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"go_service/config"
	"go_service/internal/good"
	"go_service/internal/models"
	"go_service/pkg/httpErrors"
	"go_service/pkg/utils"
	"net/http"
	"strconv"
)

type goodHandler struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
	uc     good.UseCase
}

// NewGoodHandler ..
func NewGoodHandler(cfg *config.Config, logger *zap.SugaredLogger, uc good.UseCase) goodHandler {
	return goodHandler{
		cfg:    cfg,
		logger: logger,
		uc:     uc,
	}
}

// Create
// @Summary Create Good
// @tags goods
// @description Create good
// @Param data body models.CreateGoodReq true "data for create"
// @success 200 {object} models.Good
// @router /api/v1/good/create [post]
func (h goodHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "goodHandler.Create")
		defer span.Finish()
		g := models.Good{}
		if err := c.Bind(&g); err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		goodCreated, er := h.uc.Create(ctx, &g)
		if er != nil {
			return c.JSON(httpErrors.ErrorResponse(er))
		}

		return c.JSON(http.StatusOK, goodCreated)
	}
}

// Update
// @Summary Update Good
// @tags goods
// @description Create good
// @Param data body models.UpdateGoodReq true "data for update"
// @success 200 {object} models.Good
// @router /api/v1/good/update [patch]
func (h goodHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "goodHandler.Update")
		defer span.Finish()

		g := models.Good{}
		if err := c.Bind(&g); err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		g.ID, _ = strconv.Atoi(c.Param("id"))
		g.ProjectId, _ = strconv.Atoi(c.Param("project_id"))

		goodCreated, er := h.uc.Update(ctx, &g)
		if er != nil {
			return c.JSON(httpErrors.ErrorResponse(er))
		}

		return c.JSON(http.StatusOK, goodCreated)
	}
}

// Remove
// @Summary Remove Good
// @tags goods
// @description Remove good
// @success 200 {object} models.RemoveResp
// @router /api/v1/good/remove/{id}/{project_id} [delete]
func (h goodHandler) Remove() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "goodHandler.Update")
		defer span.Finish()

		g := models.Good{}
		g.ID, _ = strconv.Atoi(c.Param("id"))
		g.ProjectId, _ = strconv.Atoi(c.Param("project_id"))

		res, er := h.uc.Remove(ctx, &g)
		if er != nil {
			return c.JSON(httpErrors.ErrorResponse(er))
		}

		return c.JSON(http.StatusOK, res)
	}
}

// List
// @Summary List Good
// @tags goods
// @description List good
// @success 200 {object} models.GoodList
// @router /api/v1/good/list/{limit}/{offset} [get]
func (h goodHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "goodHandler.Create")
		defer span.Finish()

		l := models.GoodList{}

		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		l.Limit, l.Offset = limit, offset

		list, er := h.uc.List(ctx, &l)
		if er != nil {
			return c.JSON(httpErrors.ErrorResponse(er))
		}

		return c.JSON(http.StatusOK, list)
	}
}

// Reprioritiize
// @Summary Reprioritiize Good
// @tags goods
// @description Reprioritiize good
// @Param data body models.PriorityReq true "data for reprioritiize"
// @success 200 {object} models.PriorityResp
// @router /api/v1/good/remove/{id}/{project_id} [delete]
func (h goodHandler) Reprioritiize() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "goodHandler.Update")
		defer span.Finish()
		var g models.Good
		var req models.PriorityReq

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		g.ID, _ = strconv.Atoi(c.Param("id"))
		g.ProjectId, _ = strconv.Atoi(c.Param("project_id"))
		g.Priority = req.NewPriority

		res, er := h.uc.Reprioritiize(ctx, &g)
		if er != nil {
			return c.JSON(httpErrors.ErrorResponse(er))
		}

		return c.JSON(http.StatusOK, res)
	}
}
