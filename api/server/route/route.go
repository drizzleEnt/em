package route

import (
	"github.com/drizzleent/em/api/server/middleware"
	"github.com/drizzleent/em/internal/service"
	"github.com/gin-gonic/gin"
)

func Setup(g *gin.Engine, service service.MusicService) {
	api := g.Group("/")
	api.Use(middleware.LogMiddleware)
	NewMusicRoute(api, service)
}
