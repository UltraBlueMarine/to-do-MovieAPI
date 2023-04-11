package service

import (
	server_side "server-side"
	"server-side/internal/repository"
)

type TodoItem interface {
	Create(item server_side.Movie) error
	GetAllMovies() ([]server_side.Movie, error)
	GetMovieById(id int) (server_side.Movie, error)
	UpdateMovieById(id int, movie server_side.Movie) error
	DeleteMovieById(id int) error
}

type Service struct {
	TodoItem
}

func NewService(repo *repository.Respository) *Service {
	return &Service{
		TodoItem: NewMovieService(repo),
	}
}
