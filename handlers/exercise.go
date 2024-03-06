package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type Exercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AudioPath string `json:"audio_path"`
}

type SubmitExerciseRequest struct {
	ExerciseID int    `json:"exercise_id"`
	UserID     int    `json:"user_id"`
	GenText    string `json:"generate_text"`
}

func GetExercise(db *sql.DB) http.HandlerFunc {
	// db, err := connectToDB()
	return func(w http.ResponseWriter, r *http.Request) {
		exerciseID := chi.URLParam(r, "exerciseID")
		var exercise Exercise

		query := "SELECT id, exercise_name, audio_path FROM exercises WHERE id = ?"
		row := db.QueryRow(query, exerciseID)

		err := row.Scan(&exercise.ID, &exercise.Name, &exercise.AudioPath)
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

		_, err = w.Write(exerciseJSON)
		if err != nil {
			http.Error(w, "Error writing response:", http.StatusInternalServerError)
		}
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

		_, err = w.Write(exercisesJSON)
		if err != nil {
			http.Error(w, "Error writing response:", http.StatusInternalServerError)
		}
	}
}

func SubmitExercise(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		exerciseID := chi.URLParam(r, "exerciseID")
		query := "SELECT id FROM exercises WHERE id = ?"
		row := db.QueryRow(query, exerciseID)

		var exercise Exercise
		err := row.Scan(&exercise.ID)
		if err != nil {
			http.Error(w, "the row doesn't exist in db", http.StatusInternalServerError)
			return
		}

		// Parse request body
		var reqBody SubmitExerciseRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Error decode json", http.StatusBadRequest)
			return
		}

		// TODO: Get score
		mistakes, words := countTextMistakes(reqBody.GenText)
		score := calculateScore(mistakes, words)
		res := formatFloat(score)

		// Insert data into users_results table
		_, err = db.Exec("INSERT INTO users_results (user_id, exercise_id, generate_text, result) VALUES (?, ?, ?, ?, ?)",
			reqBody.UserID, exerciseID, reqBody.GenText, res)
		if err != nil {
			http.Error(w, "Failed to insert data into users_results table", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		// Write response body
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	}
}

func countTextMistakes(genText string) (mistakes, words int) {
	return 1, 2
}

func formatFloat(value float32) string {
	return fmt.Sprintf("%.2f", value)
}

func calculateScore(mistakes, words int) float32 {
	if words == 0 {
		// Handle division by zero
		return 0.0
	}

	return float32(mistakes) / float32(words) * 100.0
}

