package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/model"
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/repository/database"
	"go.uber.org/zap"
	"net/http"
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
	if ctx.GetHeader("content-type") != "application/json" {
		ctx.JSON(http.StatusBadRequest, "only json content-type")
		return
	}
	body, err := ctx.GetRawData()
	if err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if len(body) == 0 {
		ctx.JSON(http.StatusBadRequest, "body is empty")
		return
	}
	bufRequest := model.Segment{}
	if err = json.Unmarshal(body, &bufRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	_, err = a.repo.AddSegment(bufRequest.Slug)
	if err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, "")
}

func (a *API) DeleteSegment(ctx *gin.Context) {
	if ctx.GetHeader("content-type") != "application/json" {
		ctx.JSON(http.StatusBadRequest, "only json content-type")
		return
	}
	body, err := ctx.GetRawData()
	if err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if len(body) == 0 {
		ctx.JSON(http.StatusBadRequest, "body is empty")
		return
	}
	bufRequest := model.Segment{}
	if err = json.Unmarshal(body, &bufRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err = a.repo.DeleteSegment(bufRequest.Slug)
	if err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "")
}

func (a *API) AddUserToSegments(ctx *gin.Context) {

}

func (a *API) GetUserSegments(ctx *gin.Context) {
	if ctx.GetHeader("content-type") != "application/json" {
		ctx.JSON(http.StatusBadRequest, "only json content-type")
		return
	}
	body, err := ctx.GetRawData()
	if err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if len(body) == 0 {
		ctx.JSON(http.StatusBadRequest, "body is empty")
		return
	}
	bufRequest := model.User{}
	if err = json.Unmarshal(body, &bufRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	segments, err := a.repo.GetUserSegments(bufRequest.ID)
	if err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, segments)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
