package handlers

import (
	"log"
	"net/http"
	"phonebook/models"
	"strconv"
)

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := RequireAuth(w, r)
	if !ok {
		return
	}

	users, err := models.ListUsers()
	if err != nil {
		http.Error(w, "Error listing users", http.StatusInternalServerError)
		return
	}
	LoggedUserName := GetLoggedtUserName(r)
	Tmpl.ExecuteTemplate(w, "main-layout", struct {
		LoggedUserName *string
		Data           []models.User
		TemplateName   string
	}{LoggedUserName, users, "content/users"})
}

func ShowNewUserHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := RequireAuth(w, r)
	if !ok {
		return
	}

	LoggedUserName := GetLoggedtUserName(r)
	Tmpl.ExecuteTemplate(w, "main-layout", struct {
		LoggedUserName *string
		TemplateName   string
	}{LoggedUserName, "content/users/new"})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// NOTE: For now we can store plain passwords or hash them.
	// Better practice: hash with bcrypt.
	err := models.CreateUser(username, password, 1)
	if err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func ShowEditUserHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := RequireAuth(w, r)
	if !ok {
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad id", http.StatusBadRequest)
		return
	}
	userInDb, err := models.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	LoggedUserName := GetLoggedtUserName(r)
	Tmpl.ExecuteTemplate(w, "main-layout", struct {
		LoggedUserName *string
		Data           *models.User
		TemplateName   string
	}{LoggedUserName, userInDb, "content/users/edit"})
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Bad id", http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	err = models.UpdateUser(id, username, password, 1) // updatedBy=1 dummy
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := RequireAuth(w, r)
	if !ok {
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Bad id", http.StatusBadRequest)
		return
	}
	err = models.DeleteUser(id, 1)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
