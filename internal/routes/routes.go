package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/v-vovk/femProject/internal/app"
)

func Setup(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	r.Get("/workouts/{id}", app.WorkoutHandler.GetByID)
	r.Post("/workouts", app.WorkoutHandler.Create)
	r.Put("/workouts/{id}", app.WorkoutHandler.Update)
	r.Delete("/workouts/{id}", app.WorkoutHandler.Delete)

	return r
}
