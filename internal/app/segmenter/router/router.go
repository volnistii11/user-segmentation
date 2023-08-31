package router

import (
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/api"
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/repository/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	httpServer *gin.Engine
	logger     *zap.Logger
	repo       *database.Repository
}

func New(logger *zap.Logger, repo *database.Repository) *Router {
	return &Router{
		httpServer: gin.New(),
		logger:     logger,
		repo:       repo,
	}
}

func (r *Router) Serve() *gin.Engine {
	handlers := api.New(r.logger, r.repo)

	r.httpServer.POST("/api/segment/create", handlers.CreateSegment)
	r.httpServer.POST("/api/segment/delete", handlers.DeleteSegment)
	r.httpServer.POST("/api/segment/user", handlers.UpdateUserSegments)
	r.httpServer.GET("/api/user/segment", handlers.GetUserSegments)

	return r.httpServer
}
