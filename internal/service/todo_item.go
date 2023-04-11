package service

import (
	server_side "server-side"
	"server-side/internal/repository"
)

type TodoItemService struct {
	repo repository.TodoMovie
}

func NewMovieService(repo repository.TodoMovie) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(item server_side.Movie) error {
	return s.repo.Create(item)
}

func (s *TodoItemService) GetAllMovies() ([]server_side.Movie, error) {
	return s.repo.GetAllMovies()
}

func (s *TodoItemService) GetMovieById(id int) (server_side.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s *TodoItemService) UpdateMovieById(id int, movie server_side.Movie) error {
	return s.repo.UpdateMovieById(id, movie)
}

func (s *TodoItemService) DeleteMovieById(id int) error {
	return s.repo.DeleteMovieById(id)
}
