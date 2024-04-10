package exercise

type Exercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AudioPath string `json:"audio_path"`
	Text      string `json:"text"`
}

type SubmitExerciseRequest struct {
	GenText    string `json:"gen_text"`
}
