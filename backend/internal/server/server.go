package server

import (
	"fmt"
	"net/http"

	"github.com/eulerbutcooler/go-blrplt/internal/config"
	"github.com/eulerbutcooler/go-blrplt/internal/database"
	"github.com/eulerbutcooler/go-blrplt/internal/lib/job"
	"github.com/eulerbutcooler/go-blrplt/internal/logger"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Server struct {
	Config        *config.Config
	Logger        *zerolog.Logger
	LoggerService *logger.LoggerService
	DB            *database.Database
	Redis         *redis.Client
	httpServer    *http.Server
	Job           *job.JobService
}

func New(cfg *config.Config, logger *zerolog.Logger, loggerService *logger.LoggerService) (*Server, error) {
	db, err := database.New(cfg, logger, loggerService)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
}
