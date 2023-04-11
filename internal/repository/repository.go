package repository

import (
	"database/sql"
	server_side "server-side"
)

type TodoMovie interface {
	Create(item server_side.Movie) error
	GetAllMovies() ([]server_side.Movie, error)
	GetMovieById(id int) (server_side.Movie, error)
	UpdateMovieById(id int, movie server_side.Movie) error
	DeleteMovieById(id int) error
}

type Respository struct {
	TodoMovie
}

func NewRepository(db *sql.DB) *Respository {
	return &Respository{
		TodoMovie: NewTodoItemMovie(db),
	}
}
