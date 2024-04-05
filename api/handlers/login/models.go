package login

// User model
type User struct {
	ID        int    `json:"_"`
	Sub       string `json:"_"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
