package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Batrov/go-quiz/commons"
	"github.com/Batrov/go-quiz/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("System running on %s\n", commons.PORT)

	initRouter()
}

func initRouter() {
	router := httprouter.New()
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	router.GET("/", handlers.GetIndex)
	router.GET("/survey", handlers.GetSurvey)
	router.GET("/quiz", handlers.GetQuiz)

	log.Fatal(http.ListenAndServe(commons.PORT, router))
}
