package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/v-vovk/femProject/internal/app"
)

func Setup(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		r.Get("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.GetByID))
		r.Post("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.Create))
		r.Put("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.Update))
		r.Delete("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.Delete))

	})

	r.Get("/health", app.HealthCheck)

	r.Post("/users", app.UserHandler.RegisterUser)

	r.Post("/tokens/authentication", app.TokenHandler.CreateToken)

	return r
}
