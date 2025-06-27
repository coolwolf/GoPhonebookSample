package handlers

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	LoggedUserName := GetLoggedtUserName(r)
	Tmpl.ExecuteTemplate(w, "main-layout", struct {
		LoggedUserName *string
		TemplateName   string
	}{LoggedUserName, "content/index"})
}
