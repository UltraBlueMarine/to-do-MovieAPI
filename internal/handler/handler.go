package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"server-side/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/movies", func(r chi.Router) {
			r.Get("/", h.getAllMovies)
			r.Post("/", h.createMovie)
			r.Route("/id", func(r chi.Router) {
				r.Get("/{id:[0-9]+}", h.getMovieByID)
				r.Put("/{id:[0-9]+}", h.updateMovieByID)
				r.Delete("/{id:[0-9]+}", h.deleteMovieByID)
			})
		})
	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+viper.GetString("port")+"/swagger/doc.json"),
	))

	return r
}
