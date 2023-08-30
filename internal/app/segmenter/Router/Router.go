package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Router struct {
	httpServer *gin.Engine
	logger     *zap.Logger
	db         *sqlx.DB
}

func New(logger *zap.Logger, db *sqlx.DB) *Router {
	return &Router{
		httpServer: gin.New(),
		logger:     logger,
		db:         db,
	}
}

func (r *Router) Serve() *gin.Engine {

}
