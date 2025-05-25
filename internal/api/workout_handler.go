package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/v-vovk/femProject/internal/middleware"
	"github.com/v-vovk/femProject/internal/store"
	"github.com/v-vovk/femProject/internal/utils"
)

type WorkoutHandler struct {
	workoutStore store.WorkoutStore
	logger       *log.Logger
}

func NewWorkoutHandler(workoutStore store.WorkoutStore, logger *log.Logger) *WorkoutHandler {
	return &WorkoutHandler{
		workoutStore: workoutStore,
		logger:       logger,
	}
}

func (wh *WorkoutHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	whID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: ReadIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid id parameter"})
		return
	}

	workout, err := wh.workoutStore.GetByID(whID)
	if err != nil {
		wh.logger.Printf("ERROR: GetByID: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": workout})
}

func (wh *WorkoutHandler) Create(w http.ResponseWriter, r *http.Request) {
	var workout store.Workout
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: DecodingCreateWorkout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	currentUser := middleware.GetUser(r)
	if currentUser == nil || currentUser == store.AnonymousUser {
		wh.logger.Println("ERROR: Create: user not authenticated")
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "you must be logged in to create a workout"})
		return
	}

	workout.UserID = currentUser.ID

	createdWorkout, err := wh.workoutStore.Create(&workout)
	if err != nil {
		wh.logger.Printf("ERROR: Create: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to create workout"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"workout": createdWorkout})
}

func (wh *WorkoutHandler) Update(w http.ResponseWriter, r *http.Request) {
	whID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: ReadIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid id parameter"})
		return
	}

	existingWorkout, err := wh.workoutStore.GetByID(whID)
	if err != nil {
		wh.logger.Printf("ERROR: GetByID: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	if existingWorkout == nil {
		http.NotFound(w, r)
		return
	}

	var updateWorkoutRequest struct {
		Title           *string              `json:"title"`
		Description     *string              `json:"description"`
		DurationMinutes *int                 `json:"duration_minutes"`
		CaloriesBurned  *int                 `json:"calories_burned"`
		Entries         []store.WorkoutEntry `json:"entries"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateWorkoutRequest)
	if err != nil {
		wh.logger.Printf("ERROR: DecodingUpdateWorkout: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	if updateWorkoutRequest.Title != nil {
		existingWorkout.Title = *updateWorkoutRequest.Title
	}
	if updateWorkoutRequest.Description != nil {
		existingWorkout.Description = *updateWorkoutRequest.Description
	}
	if updateWorkoutRequest.DurationMinutes != nil {
		existingWorkout.DurationMinutes = *updateWorkoutRequest.DurationMinutes
	}
	if updateWorkoutRequest.CaloriesBurned != nil {
		existingWorkout.CaloriesBurned = *updateWorkoutRequest.CaloriesBurned
	}
	if updateWorkoutRequest.Entries != nil {
		existingWorkout.Entries = updateWorkoutRequest.Entries
	}

	currentUser := middleware.GetUser(r)
	if currentUser == nil || currentUser == store.AnonymousUser {
		wh.logger.Println("ERROR: Update: user not authenticated")
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "you must be logged in to update a workout"})
		return
	}

	ownerID, err := wh.workoutStore.GetWorkoutOwner(whID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			wh.logger.Printf("ERROR: GetWorkoutOwner: %v", err)
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "workout not found"})
			return
		}

		wh.logger.Printf("ERROR: GetWorkoutOwner: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	if ownerID != currentUser.ID {
		wh.logger.Println("ERROR: Update: user not authorized to update this workout")
		utils.WriteJSON(w, http.StatusForbidden, utils.Envelope{"error": "you are not authorized to update this workout"})
		return
	}

	err = wh.workoutStore.Update(existingWorkout)
	if err != nil {
		wh.logger.Printf("ERROR: Update: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "failed to update workout"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"workout": existingWorkout})
}

func (wh *WorkoutHandler) Delete(w http.ResponseWriter, r *http.Request) {
	whID, err := utils.ReadIDParam(r)
	if err != nil {
		wh.logger.Printf("ERROR: ReadIDParam: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid id parameter"})
		return
	}

	currentUser := middleware.GetUser(r)
	if currentUser == nil || currentUser == store.AnonymousUser {
		wh.logger.Println("ERROR: Update: user not authenticated")
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "you must be logged in to update a workout"})
		return
	}

	ownerID, err := wh.workoutStore.GetWorkoutOwner(whID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			wh.logger.Printf("ERROR: GetWorkoutOwner: %v", err)
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "workout not found"})
			return
		}

		wh.logger.Printf("ERROR: GetWorkoutOwner: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	if ownerID != currentUser.ID {
		wh.logger.Println("ERROR: Update: user not authorized to delete this workout")
		utils.WriteJSON(w, http.StatusForbidden, utils.Envelope{"error": "you are not authorized to delete this workout"})
		return
	}

	err = wh.workoutStore.Delete(whID)
	if err == sql.ErrNoRows {
		wh.logger.Printf("ERROR: Delete: %v", err)
		utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "workout not found"})
		return
	}
	if err != nil {
		wh.logger.Printf("ERROR: Delete: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "error deleting workout"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
