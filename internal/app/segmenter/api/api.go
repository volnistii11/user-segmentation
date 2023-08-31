package api

import (
	"encoding/json"
	"net/http"

	"github.com/volnistii11/user-segmentation/internal/app/segmenter/model"
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/repository/database"

	"github.com/gin-gonic/gin"
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

	err = a.repo.AddSegment(bufRequest.Slug)
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

func (a *API) UpdateUserSegments(ctx *gin.Context) {
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

	bufRequest := model.UserSegments{}
	if err = json.Unmarshal(body, &bufRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := a.repo.AddUserToSegments(bufRequest.UserID, bufRequest.SegmentsToAdd); err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err := a.repo.DeleteUserFromSegments(bufRequest.UserID, bufRequest.SegmentsToDelete); err != nil {
		a.logger.Error("", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "successfully updated")
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
