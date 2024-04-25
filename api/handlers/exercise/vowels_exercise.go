package exercise

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	util "github.com/loisnicoras/handwriting-to-text/util"
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

func GetVowelsExercise(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		exerciseID := chi.URLParam(r, "exerciseID")
		var exercise VowelsExercise

		preparedQuery, err := db.Prepare("SELECT id, exercise_name, vowel, text, FROM vowels_exercises WHERE id = ?")
		if err != nil {
			log.Printf("Error preparing query: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		defer preparedQuery.Close()

		row := preparedQuery.QueryRow(exerciseID)
		err = row.Scan(&exercise.ID, &exercise.Name, &exercise.Vowel, &exercise.Text)
		if err != nil {
			log.Printf("Error retrieving exercise: %v", err)
			http.Error(w, "Exercise not found", http.StatusNotFound)
			return
		}

		exerciseJSON, err := json.Marshal(exercise)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(exerciseJSON)
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}
}

func SubmitVowelExercise(db *sql.DB, projectId, region string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Failed to retrieve session", http.StatusInternalServerError)
			return
		}

		sub, ok := session.Values["sub"].(string)
		if !ok {
			http.Error(w, "User ID not found in session", http.StatusBadRequest)
			return
		}

		exerciseID := chi.URLParam(r, "exerciseID")
		query := "SELECT id, comparison_text FROM vowels_exercises WHERE id = ?"
		row := db.QueryRow(query, exerciseID)

		var exercise VowelsExercise
		err = row.Scan(&exercise.ID, &exercise.ComparisonText)
		if err != nil {
			log.Printf("Error retrieving exercise: %v", err)
			http.Error(w, "Exercise not found", http.StatusInternalServerError)
			return
		}

		var reqBody SubmitVowelExerciseRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		score, err := util.CalculateScore(exercise.ComparisonText, reqBody.Text, projectId, region)
		if err != nil {
			log.Printf("Error calculating score: %v", err)
			http.Error(w, "Failed to calculate score", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO vowels_results (sub, exercise_id, text, result) VALUES (?, ?, ?, ?)",
			sub, exerciseID, reqBody.Text, score)
		if err != nil {
			log.Printf("Error inserting data: %v", err)
			http.Error(w, "Failed to insert data into vowels_results table", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(score); err != nil {
			log.Printf("Error encoding JSON: %v", err)
		}
	}
}
