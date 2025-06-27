package handlers

import (
	"net/http"
	"phonebook/models"
	"strconv"
)

func ListContactsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	var contacts []models.Contact
	var err error
	//logrus.Info("ListContactsHandler: query:", query)
	contacts, err = models.ListContacts(query)
	if err != nil {
		http.Error(w, "Error listing contacts", http.StatusInternalServerError)
		return
	}
	LoggedUserName := GetLoggedtUserName(r)
	//logrus.Info("ListContactsHandler: LoggedUserName:", LoggedUserName, " contacts:", contacts)
	Tmpl.ExecuteTemplate(w, "main-layout", struct {
		LoggedUserName *string
		Data           []models.Contact
		TemplateName   string
	}{LoggedUserName, contacts, "content/contacts"})
}

func ShowNewContactFormHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := RequireAuth(w, r)
	if !ok {
		return
	}
	user, _ := GetCurrentUser(r)
	LoggedUserName := GetLoggedtUserName(r)
	Tmpl.ExecuteTemplate(w, "main-layout", struct {
		LoggedUserName *string
		user           *models.User
		TemplateName   string
	}{LoggedUserName, user, "content/contacts/new"})
}

func CreateContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.FormValue("name")
	phone := r.FormValue("phone")

	// Hardcode insertedBy=1 for now
	err := models.CreateContact(name, phone, 1)
	if err != nil {
		http.Error(w, "Error creating contact", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func EditContactHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := RequireAuth(w, r)
	if !ok {
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	contact, err := models.GetContact(id)
	//log.Println("EditContactHandler: contact:", contact)
	if err != nil {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}
	user, _ := GetCurrentUser(r)
	Tmpl.ExecuteTemplate(w, "main-layout", struct {
		CurrentUser  *models.User
		Data         interface{}
		TemplateName string
	}{user, contact, "content/contacts/edit"})
}

func UpdateContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	phone := r.FormValue("phone")

	// Hardcode udpatedBy=1 for now
	err = models.UpdateContact(id, name, phone, 1)
	if err != nil {
		http.Error(w, "Error creating contact", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := RequireAuth(w, r)
	if !ok {
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	// Dummy updatedBy = 1 for now
	err = models.DeleteContact(id, 1)
	if err != nil {
		http.Error(w, "Error deleting contact", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}
