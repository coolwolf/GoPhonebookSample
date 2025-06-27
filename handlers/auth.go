package handlers

import (
	"fmt"
	"net/http"
	"phonebook/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func ShowLoginHandler(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "content/login", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	user, err := models.GetUserByUsername(username)
	//hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//hashTest := string(hashedBytes)
	//log.Printf("Login attempt for user: %s", hashTest)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Set a simple session cookie (user ID as string)
	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: fmt.Sprintf("%d", user.ID),
		Path:  "/",
	})
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "user_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetCurrentUser(r *http.Request) (*models.User, bool) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return nil, false
	}
	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return nil, false
	}
	user, err := models.GetUser(id)
	if err != nil {
		return nil, false
	}
	return user, true
}

func GetLoggedtUserName(r *http.Request) *string {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return nil
	}
	id, err := strconv.Atoi(cookie.Value)
	if err != nil {
		return nil
	}
	user, err := models.GetUser(id)
	if err != nil {
		return nil
	}
	return &user.Username
}

func RequireAuth(w http.ResponseWriter, r *http.Request) (*models.User, bool) {
	user, ok := GetCurrentUser(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return nil, false
	}
	return user, true
}
