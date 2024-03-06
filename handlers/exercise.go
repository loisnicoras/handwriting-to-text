package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type Exercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AudioPath string `json:"audio_path"`
	Text      string `json:"text"`
}

type SubmitExerciseRequest struct {
	ExerciseID int    `json:"exercise_id"`
	UserID     int    `json:"user_id"`
	GenText    string `json:"generate_text"`
}

type GeminiRequest struct {
	CorrectText   string `json:"correct_text"`
	IncorrectText string `json:"incorrect_text"`
}

type GeminiResponse struct {
	Score float64 `json:"score"`
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
		query := "SELECT id, text FROM exercises WHERE id = ?"
		row := db.QueryRow(query, exerciseID)

		var exercise Exercise
		err := row.Scan(&exercise.ID, &exercise.Text)
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

		score := calculateScore(exercise.Text, reqBody.GenText)
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

func formatFloat(value float32) string {
	return fmt.Sprintf("%.2f", value)
}

func calculateScore(correctText, genText string) float32 {
	requestPayload := GeminiRequest{
		CorrectText:   correctText,
		IncorrectText: genText,
	}

	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error encoding request payload:", err)
		return 0
	}

	// Send POST request to Gemini API
	apiUrl := "https://api.gemini.com/compare_texts" //change the api URL
	response, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error sending request to Gemini API:", err)
		return 0
	}
	defer response.Body.Close()

	fmt.Print(response)
	// Read response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	// Parse JSON response
	var geminiResponse GeminiResponse
	if err := json.Unmarshal(responseBody, &geminiResponse); err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return 0
	}
	fmt.Print(geminiResponse)

	return float32(geminiResponse.Score)
}
