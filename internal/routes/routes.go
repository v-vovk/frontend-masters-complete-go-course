package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/v-vovk/femProject/internal/app"
)

func Setup(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	return r
}
