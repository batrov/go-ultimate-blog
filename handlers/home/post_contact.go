package home

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// PostContact Handler
func PostContact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	r.ParseForm()
	log.Println("==== POST EMAIL START ====")
	log.Printf("Name: %s\n", r.FormValue("name"))
	log.Printf("Email: %s\n", r.FormValue("email"))
	log.Printf("Subject: %s\n", r.FormValue("subject"))
	log.Printf("Message: %s\n", r.FormValue("message"))
	log.Printf("Time: %v\n", time.Now())
	log.Println("==== POST EMAIL END ====")
	response := fmt.Sprintf("<p>Thanks Mr./Mrs. %s for contacting us, your message has been logged.<br>We will get back to you soon!<br><br><a href='/'>BACK</a></p>", r.FormValue("name"))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(response))

}
