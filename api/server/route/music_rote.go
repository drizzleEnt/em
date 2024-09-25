package route

import (
	"github.com/drizzleent/em/api/server/controller"
	"github.com/drizzleent/em/internal/service"
	"github.com/gin-gonic/gin"
)

func NewMusicRoute(group *gin.RouterGroup, service service.MusicService) {
	mc := controller.NewMusicController(service)

	group.GET("/library", mc.GetLibrary)
	group.GET("/text", mc.GetSong)
	group.DELETE("/delete", mc.DeleteSong)
	group.PATCH("/edit", mc.UpdateSong)
	group.POST("/add", mc.AddSong)
}
