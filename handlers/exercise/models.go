package exercise

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
