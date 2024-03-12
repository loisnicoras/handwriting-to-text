package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// User model
type User struct {
	ID        int
	Sub       string
	Email     string
	Name      string
	AvatarURL string
}

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random"
	store             = sessions.NewCookieStore([]byte("your-secret-key"))
)

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     "775860936974-jqh6iu0t16505dg53hscobepn31o8uo9.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		if state != oauthStateString {
			http.Error(w, "Invalid state parameter", http.StatusBadRequest)
			return
		}

		code := r.FormValue("code")
		token, err := googleOauthConfig.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
			return
		}

		userInfo, err := fetchGoogleUserInfo(token)
		if err != nil {
			http.Error(w, "Error fetching user info", http.StatusInternalServerError)
			return
		}

		// Check if the user is already registered
		var existingUser User
		err = db.QueryRow("SELECT id, sub, email, name, avatar_url FROM users WHERE sub=?", userInfo["sub"]).Scan(
			&existingUser.ID,
			&existingUser.Sub,
			&existingUser.Email,
			&existingUser.Name,
			&existingUser.AvatarURL,
		)

		if err == sql.ErrNoRows {
			// If the user is not registered, insert a new user
			_, err = db.Exec("INSERT INTO users (sub, email, name, avatar_url) VALUES (?, ?, ?, ?)",
				userInfo["sub"], userInfo["email"], userInfo["name"], userInfo["picture"])
			if err != nil {
				http.Error(w, "Error creating new user", http.StatusInternalServerError)
				return
			}
		} else if err != nil {
			http.Error(w, "Error checking existing user", http.StatusInternalServerError)
			return
		}

		// Set user info in session (you can replace this with a JWT or any other session mechanism)
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Error creating session", http.StatusInternalServerError)
			return
		}

		session.Values["userID"] = userInfo["sub"].(string)
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isUserLoggedIn(r) {
			// http.Redirect(w, r, "/login", http.StatusSeeOther)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}

func isUserLoggedIn(r *http.Request) bool {
	session, err := store.Get(r, "session-name")
	if err != nil {
		// Handle the error as needed, e.g., log it or return false
		return false
	}

	userID, ok := session.Values["userID"].(string)
	fmt.Println(userID)

	return ok && userID != ""
}

func fetchGoogleUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
