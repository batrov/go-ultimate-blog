package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	go_quiz "github.com/Batrov/go-ultimate-blog/handlers/go-quiz"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("System running on %s\n", commons.PORT)

	initRouter()
}

func initRouter() {
	router := httprouter.New()
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	router.GET("/go-quiz", go_quiz.GetIndex)
	router.GET("/go-quiz/survey", go_quiz.GetSurvey)
	router.GET("/go-quiz/quiz", go_quiz.GetQuiz)

	log.Fatal(http.ListenAndServe(commons.PORT, router))
}
