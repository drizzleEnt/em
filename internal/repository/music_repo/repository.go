package musicrepo

import (
	"context"

	"github.com/drizzleent/em/internal/model"
	"github.com/drizzleent/em/pkg/client/db"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) Add(context.Context, *model.Music) error {
	return nil
}
func (r *repo) Delete(context.Context) error {
	return nil
}
func (r *repo) Update(context.Context) error {
	return nil
}
func (r *repo) GetLibrary(context.Context) error {
	return nil
}
func (r *repo) GetSong(context.Context) error {
	return nil
}
