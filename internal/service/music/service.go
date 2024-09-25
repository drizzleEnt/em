package music

type musicService struct {
}

func NewMusicService() *musicService {
	return &musicService{}
}

func (s *musicService) Add() error {
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
