package api

import (
	"github.com/gin-gonic/gin"
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/repository/database"
	"go.uber.org/zap"
)

type API struct {
	logger *zap.Logger
	repo   *database.Repository
}

func New(logger *zap.Logger, repo *database.Repository) *API {
	return &API{
		logger: logger,
		repo:   repo,
	}
}

func (a *API) CreateSegment(ctx *gin.Context) {

}

func (a *API) DeleteSegment(ctx *gin.Context) {

}

func (a *API) AddUserToSegments(ctx *gin.Context) {

}

func (a *API) GetUserSegments(ctx *gin.Context) {

}
