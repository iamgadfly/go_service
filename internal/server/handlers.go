package server

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	goodHttp "go_service/internal/good/delivery/http"
	goodRepository "go_service/internal/good/repository"
	goodUseCase "go_service/internal/good/usecase"
	projectHttp "go_service/internal/project/delivery/http"
	projectRepository "go_service/internal/project/repository"
	projectUseCase "go_service/internal/project/usecase"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := e.Group("/api/v1")

	// init good
	goodsGroup := v1.Group("/good")
	gRepo := goodRepository.NewGoodRepo(s.logger, s.db)
	gRedisRepo := goodRepository.NewGoodRedisRepository(s.logger, s.redis)
	goodUC := goodUseCase.NewGoodUC(s.cfg, s.logger, gRedisRepo, gRepo)
	goodHandler := goodHttp.NewGoodHandler(s.cfg, s.logger, goodUC)
	goodHttp.MapProductRoutes(goodsGroup, goodHandler)

	// init project
	projectsGroup := v1.Group("/project")
	pRepo := projectRepository.NewProjectRepo(s.db, s.logger)
	projectUC := projectUseCase.NewProjectUC(s.cfg, s.logger, pRepo)
	projectHandler := projectHttp.NewProjectHandler(s.cfg, s.logger, projectUC)
	projectHttp.MapProductRoutes(projectsGroup, projectHandler)

	return nil
}
