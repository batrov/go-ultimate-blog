package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	goquiz "github.com/Batrov/go-ultimate-blog/handlers/goquiz"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("System running on %s\n", commons.PORT)

	initRouter()
}

func initRouter() {
	router := httprouter.New()
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	router.GET("/", goquiz.GetIndex)
	router.GET("/goquiz", goquiz.GetIndex)
	router.GET("/goquiz/survey", goquiz.GetSurvey)
	router.GET("/goquiz/quiz", goquiz.GetQuiz)

	log.Fatal(http.ListenAndServe(commons.PORT, router))
}
