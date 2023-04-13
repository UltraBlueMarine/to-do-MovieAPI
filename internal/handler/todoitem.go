package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	server_side "server-side"
	"strconv"
)

// ShowAccount godoc
// @Summary      Show an account
// @Description  create movie
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]

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

	err = h.service.Create(movie)
}

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
