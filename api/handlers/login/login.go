package login

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random"
	userInfo          = "https://www.googleapis.com/oauth2/v3/userinfo"
	store             = sessions.NewCookieStore([]byte("your-secret-key"))
	webURL            = "http://localhost:3000"
)

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     "775860936974-jqh6iu0t16505dg53hscobepn31o8uo9.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	store.Options.MaxAge = 36000
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
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
			log.Println("Error fetching user info:", err)
			http.Error(w, "Error fetching user info", http.StatusInternalServerError)
			return
		}

		err = handleUser(db, userInfo)
		if err != nil {
			http.Error(w, "Error handling user", http.StatusInternalServerError)
			log.Println("Error handling user:", err)
			return
		}

		setSession(w, r, userInfo)

		http.Redirect(w, r, webURL, http.StatusTemporaryRedirect)
	}
}

// LogOut handles user logout
func LogOut(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		clearSession(w, r)

		http.Redirect(w, r, webURL, http.StatusSeeOther)
	}
}

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if !isUserLoggedIn(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetUserData(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		if !isUserLoggedIn(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userInfo, err := getUserInfoFromSession(r)
		if err != nil {
			http.Error(w, "Failed to retrieve user info from session", http.StatusInternalServerError)
			log.Println("Failed to retrieve user info from session:", err)
			return
		}

		userData, err := getUserDataFromDB(db, userInfo["sub"].(string))
		if err != nil {
			http.Error(w, "Failed to retrieve user data from database", http.StatusInternalServerError)
			log.Println("Failed to retrieve user data from database:", err)
			return
		}

		jsonData, err := json.Marshal(userData)
		if err != nil {
			http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
			log.Println("Failed to marshal JSON response:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func fetchGoogleUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get(userInfo)
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

// handleUser checks if the user is registered and inserts if not
func handleUser(db *sql.DB, userInfo map[string]interface{}) error {
	// Check if the user is already registered
	var existingUser User
	err := db.QueryRow("SELECT id, sub, email, name, avatar_url FROM users WHERE sub=?", userInfo["sub"]).Scan(
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
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

// setSession sets user info in session
func setSession(w http.ResponseWriter, r *http.Request, userInfo map[string]interface{}) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("Error creating session:", err)
		return
	}

	session.Values["sub"] = userInfo["sub"].(string)

	err = session.Save(r, w)
	if err != nil {
		log.Println("Error saving session:", err)
	}
}

func clearSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		log.Println("Error retrieving session:", err)
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		log.Println("Error saving session:", err)
	}
}

func isUserLoggedIn(r *http.Request) bool {
	session, err := store.Get(r, "session-name")
	if err != nil {
		// Handle the error as needed, e.g., log it or return false
		return false
	}

	userID, ok := session.Values["sub"].(string)
	return ok && userID != ""
}

// getUserInfoFromSession retrieves user info from session
func getUserInfoFromSession(r *http.Request) (map[string]interface{}, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return nil, err
	}

	userInfo, ok := session.Values["sub"].(string)
	if !ok || userInfo == "" {
		return nil, errors.New("user ID not found in session")
	}

	return map[string]interface{}{"sub": userInfo}, nil
}

// getUserDataFromDB retrieves user data from the database
func getUserDataFromDB(db *sql.DB, sub string) (User, error) {
	var user User
	query := "SELECT name, email, avatar_url FROM users WHERE sub = ?"
	err := db.QueryRow(query, sub).Scan(&user.Name, &user.Email, &user.AvatarURL)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
