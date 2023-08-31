package main

import (
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/repository/database"
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/router"
	"github.com/volnistii11/user-segmentation/internal/config"
	"go.uber.org/zap"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.New()
	err := cfg.Parse()
	if err != nil {
		log.Fatalf("cannot parse config: %s", err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to start logger")
	}
	defer logger.Sync()

	conn, err := database.NewConnection(cfg.DatabaseDriver, cfg.DatabaseDSN)
	if err != nil {
		logger.Error("failed to create db connection", zap.Error(err))
	}

	repo := database.New(conn)

	router := router.New(logger, repo)
	if err = router.Serve().Run(cfg.HTTPServerAddress); err != nil {
		logger.Error("failed to start http server", zap.Error(err))
	}
}
