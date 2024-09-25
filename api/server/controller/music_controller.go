package controller

import (
	"net/http"

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
	err := mc.srv.Add()
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "failed to add song: "+err.Error())
	}
}

func (mc *musicController) DeleteSong(ctx *gin.Context) {

}

func (mc *musicController) UpdateSong(ctx *gin.Context) {
}
