package music

import (
	"context"

	"github.com/drizzleent/em/internal/model"
)

type musicService struct {
}

func NewMusicService() *musicService {
	return &musicService{}
}

func (s *musicService) Add(ctx context.Context, info *model.Music) error {
	// request for info
	// save in db
	return nil
}
func (s *musicService) DeleteSong() error {
	return nil
}
func (s *musicService) UpdateSong() error {
	return nil
}
func (s *musicService) GetLibrary() error {
	return nil
}
func (s *musicService) GetSong() error {
	return nil
}
