package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"io"
	"net/http"
	server_side "server-side"

	"strconv"
)

// POST Movie
// @tags POST Movie
// @Summary create movie
// @Description create movie
// @Accept  json
// @Produce  json
// @Success 200 "No Content"
// @Failure 404 "No Content"
// @Router /movies  [POST]
func (h *Handler) createMovie(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var movie server_side.Movie
	if err = json.Unmarshal(reqBytes, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = h.service.Create(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}

// GET All Movies
// @tags GET All Movies
// @Summary show all movies
// @Description show all movies
// @Accept  json
// @Produce  json
// @Success 200 {array} server_side.Movie
// @Failure 404 {object} server_side.Movie
// @Router /movies  [get]
func (h *Handler) getAllMovies(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/json")

	tmp, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(tmp)
}

// GET Movie by Id
// @tags GET Movie by Id
// @Summary show movie by Id
// @Description show movie by Id
// @Accept  json
// @Produce  json
// @Success 200 {object} server_side.Movie
// @Failure 404 {object} server_side.Movie
// @Router /movies/id/{id}  [get]
func (h *Handler) getMovieByID(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item, err := h.service.GetMovieById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmp, err := json.Marshal(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(tmp)
}

// PUT Movie by Id
// @tags UPDATE Movie by Id
// @Summary update movie by Id
// @Description update movie by Id
// @Accept  json
// @Produce  json
// @Success 200 "No Content"
// @Failure 404 "No Content"
// @Router /movies/id/{id}  [PUT]
func (h *Handler) updateMovieByID(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reqBytes, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var movie server_side.Movie

	if err := json.Unmarshal(reqBytes, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = h.service.UpdateMovieById(id, movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// DELETE Movie by Id
// @tags DELETE Movie by Id
// @Summary delete movie by Id
// @Description delete movie by Id
// @Accept  json
// @Produce  json
// @Success 200 "No Content"
// @Failure 404 "No Content"
// @Router /movies/id/{id}  [delete]
func (h *Handler) deleteMovieByID(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = h.service.DeleteMovieById(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
