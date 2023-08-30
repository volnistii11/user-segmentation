package segmenter

import (
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/repository/database"
	"github.com/volnistii11/user-segmentation/internal/config"
	"go.uber.org/zap"
	"log"
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

	db, err := database.NewConnection(cfg.DatabaseDriver, cfg.DatabaseDSN)

}
