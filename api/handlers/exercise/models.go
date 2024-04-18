package exercise

type AudioExercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	AudioPath string `json:"audio_path"`
	Text      string `json:"text"`
}

type SubmitAudioExerciseRequest struct {
	GenText    string `json:"gen_text"`
}

type VowelsExercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Text      string `json:"text"`
}