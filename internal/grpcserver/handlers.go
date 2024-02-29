package grpcserver

import (
	repository2 "go_service/internal/good/repository"
	goodUseCase "go_service/internal/good/usecase"
	"go_service/internal/project/repository"
	projectUseCase "go_service/internal/project/usecase"
	"go_service/pkg/api/good"
	good2 "go_service/pkg/api/good/proto"
	"go_service/pkg/api/project"
	proto "go_service/pkg/api/project/proto"
	"google.golang.org/grpc"
)

// MapServices ..
func (s *ServerGRPC) MapServices(server *grpc.Server) {
	projectRepo := repository.NewProjectRepo(s.db, s.logger)
	projectUC := projectUseCase.NewProjectUC(s.cfg, s.logger, projectRepo)
	sProject := project.NewProjectServer(s.cfg, projectRepo, projectUC)

	goodRepo := repository2.NewGoodRepo(s.logger, s.db)
	goodRedisRepo := repository2.NewGoodRedisRepository(s.logger, s.redis)
	goodUC := goodUseCase.NewGoodUC(s.cfg, s.logger, goodRedisRepo, goodRepo)
	sGood := good.NewGoodServer(s.cfg, goodRepo, goodUC)

	proto.RegisterProjectServer(server, sProject)
	good2.RegisterGoodServer(server, sGood)
}
