package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/vertexai/genai"
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

type Candidate struct {
	Index   int `json:"Index"`
	Content struct {
		Role  string   `json:"Role"`
		Parts []string `json:"Parts"`
	} `json:"Content"`
	FinishReason     int            `json:"FinishReason"`
	SafetyRatings    []SafetyRating `json:"SafetyRatings"`
	FinishMessage    string         `json:"FinishMessage"`
	CitationMetadata interface{}    `json:"CitationMetadata"`
}

type SafetyRating struct {
	Category    int  `json:"Category"`
	Probability int  `json:"Probability"`
	Blocked     bool `json:"Blocked"`
}

func GetExercise(db *sql.DB) http.HandlerFunc {
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

func SubmitExercise(db *sql.DB, projectId, region string) http.HandlerFunc {
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

		score, err := calculateScore(exercise.Text, reqBody.GenText, projectId, region)
		if err != nil {
			http.Error(w, "Failed to calculate the score", http.StatusInternalServerError)
			return
		}

		// Insert data into users_results table
		_, err = db.Exec("INSERT INTO users_results (user_id, exercise_id, generate_text, result) VALUES (?, ?, ?, ?, ?)",
			reqBody.UserID, exerciseID, reqBody.GenText, score)
		if err != nil {
			http.Error(w, "Failed to insert data into users_results table", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		// Write response body
		if err := json.NewEncoder(w).Encode(score); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
			return
		}
	}
}

func calculateScore(correctText, genText, projectId, region string) (int, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectId, region)
	if err != nil {
		return 0, fmt.Errorf("Failed create new client: %w", err)
	}
	gemini := client.GenerativeModel("gemini-pro-vision")

	prompt := genai.Text("Can you give me a score (just the score no more words) that is between 0-100 from comparing the first text with the second. First is the correct text. The correct text is: " + correctText + " The incorrect text is " + genText)
	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return 0, fmt.Errorf("Failed to generate content: %w", err)
	}
	rb, _ := json.MarshalIndent(resp, "", "  ")

	type Response struct {
		Candidates []Candidate `json:"Candidates"`
	}

	// Unmarshal the JSON response string into the Response struct
	var response Response
	err = json.Unmarshal([]byte(rb), &response)
	if err != nil {
		return 0, fmt.Errorf("Failed to unmarshal the json: %w", err)
	}

	// Access the "Parts" data from the first candidate
	parts := response.Candidates[0].Content.Parts

	// Convert the string value to float64
	floatValue, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("Failed to parse float value: %w", err)
	}

	// Convert float64 to integer
	intValue := int(floatValue)

	return intValue, nil
}
