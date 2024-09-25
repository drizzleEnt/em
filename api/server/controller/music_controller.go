package controller

import (
	"net/http"

	"github.com/drizzleent/em/internal/converter"
	"github.com/drizzleent/em/internal/logger"
	"github.com/drizzleent/em/internal/service"
	"github.com/gin-gonic/gin"
)

type musicController struct {
	srv service.MusicService
}

func NewMusicController(srv service.MusicService) *musicController {
	return &musicController{
		srv: srv,
	}
}

func (mc *musicController) GetLibrary(ctx *gin.Context) {

}

func (mc *musicController) GetSong(ctx *gin.Context) {
}

func (mc *musicController) AddSong(ctx *gin.Context) {
	info, err := converter.FromReqToSong(ctx)
	if err != nil {
		logger.Error("failed to parse song: " + err.Error())
		NewErrorResponse(ctx, http.StatusBadRequest, "failed to parse song: "+err.Error())
		return
	}
	err = mc.srv.Add(ctx, info)
	if err != nil {
		logger.Error("failed to add song: " + err.Error())
		NewErrorResponse(ctx, http.StatusInternalServerError, "failed to add song: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "song was added")
}

func (mc *musicController) DeleteSong(ctx *gin.Context) {

}

func (mc *musicController) UpdateSong(ctx *gin.Context) {
}
