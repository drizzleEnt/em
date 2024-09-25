package service

import (
	"context"

	"github.com/drizzleent/em/internal/model"
)

type MusicService interface {
	Add(context.Context, *model.Music) error
	DeleteSong() error
	UpdateSong() error
	GetLibrary() error
	GetSong() error
}
