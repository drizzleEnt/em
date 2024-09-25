package repository

import (
	"context"

	"github.com/drizzleent/em/internal/model"
)

type MusicRepository interface {
	Add(context.Context, *model.Music) error
	Delete(context.Context) error
	Update(context.Context) error
	GetLibrary(context.Context) error
	GetSong(context.Context) error
}
