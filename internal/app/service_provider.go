package app

import (
	"context"
	"log"

	"github.com/drizzleent/em/api/server/route"
	"github.com/drizzleent/em/internal/config"
	"github.com/drizzleent/em/internal/config/env"
	"github.com/drizzleent/em/internal/service"
	"github.com/drizzleent/em/internal/service/music"
	"github.com/drizzleent/em/pkg/client/db"
	"github.com/drizzleent/em/pkg/client/db/pg"
	"github.com/gin-gonic/gin"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig

	dbClient db.Client

	musicService service.MusicService

	engine *gin.Engine
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if nil == s.httpConfig {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}
		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PGConfig() config.HTTPConfig {
	if nil == s.pgConfig {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if nil == s.dbClient {
		cl, err := pg.New(ctx, s.PGConfig().Address())
		if err != nil {
			log.Fatalf("Failed to create db client %s", err.Error())
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("Failed to ping db %s", err.Error())
		}
		s.dbClient = cl
	}
	return s.dbClient
}



func (s *serviceProvider) MusicService(ctx context.Context) service.MusicService {
	if nil == s.musicService {
		s.musicService = music.NewMusicService()
	}
	return s.musicService
}

func (s *serviceProvider) Engine(ctx context.Context) *gin.Engine {
	if nil == s.engine {
		e := gin.New()
		route.Setup(e, s.MusicService(ctx))
		s.engine = e
	}

	return s.engine
}
