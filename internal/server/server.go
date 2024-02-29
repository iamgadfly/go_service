package server

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go_service/config"
	"log"
	"time"
)

const (
	ctxTimeout = 5
)

// Server struct
type Server struct {
	cfg        *config.Config
	redis      *redis.Client
	port       string
	echo       *echo.Echo
	db         *sqlx.DB
	clickHouse *sql.DB
	logger     *zap.SugaredLogger
}

//clickHouse *sql.DB,

// NewServer New Server constructor
func NewServer(cfg *config.Config, redis *redis.Client, port string, db *sqlx.DB, logger *zap.SugaredLogger) *Server {
	return &Server{
		cfg:   cfg,
		redis: redis,
		echo:  echo.New(),
		db:    db,
		port:  port,
		//clickHouse: clickHouse,
		logger: logger,
	}
}

func (s *Server) Run(er chan error) {
	log.Println("staring http server!")
	if err := s.MapHandlers(s.echo); err != nil {
		er <- err
	}
	if err := s.echo.Start(s.port); err != nil {
		er <- err
	}

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		er <- err
	}

	er <- nil
}
