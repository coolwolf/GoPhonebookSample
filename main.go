package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"phonebook/db"
	"phonebook/handlers"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logrus.Infof("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		logrus.Infof("Completed in %v", time.Since(start))
	})
}

func init() {
	funcMap := template.FuncMap{
		"currentYear": func() int {
			return time.Now().Year()
		},
		"yield": func(name string, data interface{}) (template.HTML, error) {
			var buf bytes.Buffer
			err := handlers.Tmpl.ExecuteTemplate(&buf, name, data)
			return template.HTML(buf.String()), err
		},
	}
	var err error
	handlers.Tmpl, err = template.New("").Funcs(funcMap).ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("APP_ENV")
	logrus.Info("Starting Phonebook Application", " Environment:", env)
	if env == "development" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
	mux := http.NewServeMux()

	db.ConnectDB("phonebook.db")
	db.CreateTables()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	mux.HandleFunc("/", handlers.ListContactsHandler)

	mux.HandleFunc("/contacts", handlers.ListContactsHandler)
	mux.HandleFunc("/contacts/new", handlers.ShowNewContactFormHandler)
	mux.HandleFunc("/contacts/create", handlers.CreateContactHandler) //POST
	mux.HandleFunc("/contacts/edit", handlers.EditContactHandler)     // GET
	mux.HandleFunc("/contacts/update", handlers.UpdateContactHandler) // POST
	mux.HandleFunc("/contacts/delete", handlers.DeleteContactHandler) // BACKEND

	mux.HandleFunc("/users", handlers.ListUsersHandler)
	mux.HandleFunc("/users/new", handlers.ShowNewUserHandler)
	mux.HandleFunc("/users/create", handlers.CreateUserHandler) // POST
	mux.HandleFunc("/users/edit", handlers.ShowEditUserHandler) //GET
	mux.HandleFunc("/users/update", handlers.UpdateUserHandler) // POST
	mux.HandleFunc("/users/delete", handlers.DeleteUserHandler) // BACKEND

	mux.HandleFunc("/login", handlers.ShowLoginHandler) // GET
	mux.HandleFunc("/dologin", handlers.LoginHandler)   // POST
	mux.HandleFunc("/logout", handlers.LogoutHandler)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
