package repository

import (
	"database/sql"
	"errors"
	"fmt"
	server_side "server-side"
)

type TodoItemPostgres struct {
	db *sql.DB
}

func NewTodoItemMovie(db *sql.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(movie server_side.Movie) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if r.ItemExists(movie) {
		return errors.New("Item alredy exists!")
	}

	_, err = tx.Exec("insert into movies1 (title, year, director, done) values ($1, $2, $3, $4)", movie.Title, movie.Year, movie.Director, movie.Done)

	if err != nil {
		defer tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		defer tx.Rollback()
		return err
	}

	return nil
}

func (r *TodoItemPostgres) GetAllMovies() ([]server_side.Movie, error) {
	rows, err := r.db.Query("select * from movies1")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	fmt.Println(rows)
	movies := make([]server_side.Movie, 0)
	for rows.Next() {
		m := server_side.Movie{}
		err := rows.Scan(&m.Id, &m.Title, &m.Year, &m.Director, &m.Done)
		if err != nil {

			return nil, err
		}
		movies = append(movies, m)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return movies, err
}

func (r *TodoItemPostgres) GetMovieById(id int) (server_side.Movie, error) {
	var movie server_side.Movie
	row := r.db.QueryRow("select * from movies1 where id = $1", id).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Director, &movie.Done)
	if row != nil {
		if errors.Is(row, sql.ErrNoRows) {
			return server_side.Movie{}, row
		}
	}
	return movie, nil
}

func (r *TodoItemPostgres) UpdateMovieById(id int, movie server_side.Movie) error {
	tx, err := r.db.Begin()
	if err != nil {
		defer tx.Rollback()
		return err
	}

	_, err = tx.Exec("update movies1 set title=$1, year=$2, director=$3, done=$4 where id=$5", movie.Title, movie.Year, movie.Director, movie.Done, id)
	if err != nil {
		defer tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		defer tx.Rollback()
		return err
	}

	return nil
}

func (r *TodoItemPostgres) DeleteMovieById(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		defer tx.Rollback()
		return err
	}

	_, err = tx.Exec("delete from movies1 where id=$1", id)
	if err != nil {
		defer tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		defer tx.Rollback()
		return err
	}

	return nil
}

func (r *TodoItemPostgres) ItemExists(item server_side.Movie) bool {
	var exists int
	err := r.db.QueryRow("select id from movies1 where title=$1 and director=$2", item.Title, item.Director).Scan(&exists)
	if err != nil {
		return false
	}
	return true
}
