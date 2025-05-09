package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	paramsWhID := chi.URLParam(r, "id")
	if paramsWhID == "" {
		http.NotFound(w, r)
		return
	}

	whID, err := strconv.ParseInt(paramsWhID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Workout with ID %d\n", whID)
}

func (wh *WorkoutHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new workout\n")
}
