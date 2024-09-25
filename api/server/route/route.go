package route

import (
	"github.com/drizzleent/em/internal/service"
	"github.com/gin-gonic/gin"
)

func Setup(g *gin.Engine, service service.MusicService) {
	api := g.Group("/")
	NewMusicRoute(api, service)
}
