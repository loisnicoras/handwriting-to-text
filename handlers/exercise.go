package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type Exercise struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Audio_path    string `json:"audio_path"`
}

func GetExercise(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDB()
	exerciseID := chi.URLParam(r, "exerciseID")

	if err != nil {
		http.Error(w, "failed to connect to the db", http.StatusInternalServerError)
		return
	}

	var exercise Exercise

	query := "SELECT * FROM your_table WHERE your_id = ?"
	row := db.QueryRow(query, exerciseID)

	err = row.Scan(&exercise.ID, &exercise.Name, &exercise.Audio_path)
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

func GetExercises(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDB()
	if err != nil {
		http.Error(w, "failed to connect to the db", http.StatusInternalServerError)
		return
	}

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

	// Check for errors during row iteration
	// if err = rows.Err(); err != nil {
	// 	http.Error(w, "failed to connect to the db", http.StatusInternalServerError)
	// 	return
	// }

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

func SubmitExercise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
