package exercise

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	util "github.com/loisnicoras/handwriting-to-text/util"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetExercise(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		exerciseID := chi.URLParam(r, "exerciseID")
		var exercise Exercise

		query := "SELECT id, exercise_name, audio_path FROM exercises WHERE id = ?"
		row := db.QueryRow(query, exerciseID)

		err := row.Scan(&exercise.ID, &exercise.Name, &exercise.AudioPath)
		if err != nil {
			log.Printf("Error retrieving exercise: %v", err)
			http.Error(w, "Exercise not found", http.StatusInternalServerError)
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

func GetExercises(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		rows, err := db.Query("SELECT id, exercise_name FROM exercises")
		if err != nil {
			log.Printf("Error retrieving exercises: %v", err)
			http.Error(w, "Failed to get exercises", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var exercises []Exercise

		for rows.Next() {
			var exercise Exercise
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

func SubmitExercise(db *sql.DB, projectId, region string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		exerciseID := chi.URLParam(r, "exerciseID")
		query := "SELECT id, text FROM exercises WHERE id = ?"
		row := db.QueryRow(query, exerciseID)

		var exercise Exercise
		err := row.Scan(&exercise.ID, &exercise.Text)
		if err != nil {
			log.Printf("Error retrieving exercise: %v", err)
            http.Error(w, "Exercise not found", http.StatusInternalServerError)
			return
		}

		var reqBody SubmitExerciseRequest
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			log.Printf("Error decoding JSON: %v", err)
            http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		score, err := util.CalculateScore(exercise.Text, reqBody.GenText, projectId, region)
		if err != nil {
			log.Printf("Error calculating score: %v", err)
			http.Error(w, "Failed to calculate score", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO users_results (user_id, exercise_id, photo_text, generate_text, result) VALUES (?, ?, ?, ?, ?)",
			reqBody.UserID, exerciseID, "", reqBody.GenText, score)
		if err != nil {
			log.Printf("Error inserting data: %v", err)
			http.Error(w, "Failed to insert data into users_results table", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(score); err != nil {
            log.Printf("Error encoding JSON: %v", err)
		}
	}
}
