package http

import "github.com/labstack/echo/v4"

func MapProductRoutes(e *echo.Group, h goodHandler) {
	e.POST("/create", h.Create())
	e.PATCH("/update/:id/:project_id", h.Update())
	e.DELETE("/remove/:id/:project_id", h.Remove())
	e.GET("/list", h.List())
	e.PATCH("/reprioritiize/:id/:project_id", h.Reprioritiize())
}
