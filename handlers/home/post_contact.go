package home

import (
	"fmt"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/services"
	"github.com/julienschmidt/httprouter"
)

// PostContact Handler
func PostContact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	err := services.GetService().HomeService.PostContact(commons.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	})
	response := fmt.Sprintf("<p>Thanks Mr./Mrs. %s for contacting us, your message has been logged.<br>We will get back to you soon!<br><br><a href='/'>BACK</a></p>", r.FormValue("name"))
	if err != nil {
		response = err.Error()
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(response))
}
