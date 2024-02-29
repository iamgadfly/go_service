package grpcserver

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go_service/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

// ServerGRPC ..
type ServerGRPC struct {
	db     *sqlx.DB
	redis  *redis.Client
	cfg    *config.Config
	logger *zap.SugaredLogger
}

// NewServerGRPC ..
func NewServerGRPC(db *sqlx.DB, redis *redis.Client, cfg *config.Config, logger *zap.SugaredLogger) *ServerGRPC {
	return &ServerGRPC{
		db:     db,
		redis:  redis,
		cfg:    cfg,
		logger: logger,
	}
}

// Run ..
func (s *ServerGRPC) Run(er chan error) {
	log.Println("starting grpc server!")

	server := grpc.NewServer()
	s.MapServices(server)

	l, err := net.Listen("tcp", s.cfg.Server.PprofPort)
	if err != nil {
		er <- err
	}

	defer l.Close()
	defer server.Stop()

	err = server.Serve(l)
	if err != nil {
		er <- err
	}

	er <- nil
}
