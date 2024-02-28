package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type Exercise struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Audio_path string `json:"audio_path"`
}

func GetExercise(db *sql.DB) http.HandlerFunc {
	// db, err := connectToDB()
	return func(w http.ResponseWriter, r *http.Request) {
		exerciseID := chi.URLParam(r, "exerciseID")
		var exercise Exercise

		query := "SELECT id, exercise_name, audio_path FROM exercises WHERE id = ?"
		row := db.QueryRow(query, exerciseID)

		err := row.Scan(&exercise.ID, &exercise.Name, &exercise.Audio_path)
		if err != nil {
			http.Error(w, "the row doesn't exist in db", http.StatusInternalServerError)
			return
		}

		exerciseJSON, err := json.Marshal(exercise)
		if err != nil {
			http.Error(w, "failed to marshal the json", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(exerciseJSON)
	}
}

func GetExercises(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, exercise_name FROM exercises")
		if err != nil {
			http.Error(w, "failed to get data from exercises table", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create a slice to hold the exercises
		var exercises []Exercise

		// Iterate over the result set and append exercises to the slice
		for rows.Next() {
			var exercise Exercise
			err := rows.Scan(&exercise.ID, &exercise.Name)
			if err != nil {
				http.Error(w, "failed to scan the rows", http.StatusInternalServerError)
				return
			}
			exercises = append(exercises, exercise)
		}

		// Marshal exercises slice to JSON
		exercisesJSON, err := json.Marshal(exercises)
		if err != nil {
			http.Error(w, "failed to marshal the json", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(exercisesJSON)
	}
}

func SubmitExercise(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}
