package exercise

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func GetVowelsExercises(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		rows, err := db.Query("SELECT id, exercise_name FROM vowels_exercises")
		if err != nil {
			log.Printf("Error retrieving vowels_exercises: %v", err)
			http.Error(w, "Failed to get vowels_exercises", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var exercises []VowelsExercise

		for rows.Next() {
			var exercise VowelsExercise
			err := rows.Scan(&exercise.ID, &exercise.Name)
			if err != nil {
				log.Printf("Error scanning row: %v", err)
				http.Error(w, "Failed to scan rows", http.StatusInternalServerError)
				return
			}
			exercises = append(exercises, exercise)
		}

		exercisesJSON, err := json.Marshal(exercises)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(exercisesJSON)
		if err != nil {
			http.Error(w, "Error writing response:", http.StatusInternalServerError)
		}
	}
}
