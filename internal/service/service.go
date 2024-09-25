package service

type MusicService interface {
	Add() error
	DeleteSong() error
	UpdateSong() error
	GetLibrary() error
	GetSong() error
}
