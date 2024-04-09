package exercise

type Exercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AudioPath string `json:"audio_path"`
	Text      string `json:"text"`
}

type SubmitExerciseRequest struct {
	UserID     int    `json:"user_id"`
	GenText    string `json:"gen_text"`
}
