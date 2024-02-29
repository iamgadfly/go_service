package http

import "github.com/labstack/echo/v4"

func MapProductRoutes(e *echo.Group, h projectHandler) {
	e.POST("/create", h.Create())
}
