package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/handlers/goquiz"
	"github.com/Batrov/go-ultimate-blog/handlers/home"
	"github.com/Batrov/go-ultimate-blog/handlers/iganalyzer"
	"github.com/Batrov/go-ultimate-blog/handlers/openmap"
	"github.com/Batrov/go-ultimate-blog/handlers/retirementcalc"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Printf("System running on %s\n", commons.GetPort())
	initRouter()
}

func initRouter() {
	router := httprouter.New()
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	router.GET("/", home.GetIndex)
	router.GET("/about", home.GetAbout)

	// Go Quiz
	router.GET("/goquiz", goquiz.GetIndex)
	router.GET("/goquiz/survey", goquiz.GetSurvey)
	router.POST("/goquiz/survey", goquiz.PostSurvey)
	router.GET("/goquiz/quiz", goquiz.GetQuiz)
	router.POST("/goquiz/reset-answer", goquiz.PostResetAnswer)
	router.GET("/goquiz/fetch-answer", goquiz.GetFetchAnswer)

	// Retirement Calculator
	router.GET("/retirement-calculator", retirementcalc.GetIndex)
	router.POST("/retirement-calculator", retirementcalc.PostStatistics)
	router.POST("/retirement-calculator/statistics", retirementcalc.PostStatistics)
	router.GET("/retirement-calculator/statistics", retirementcalc.GetStatistics)

	// Insta Analyzer
	router.GET("/instagram-analyzer", iganalyzer.GetIndex)

	// Open Map
	router.GET("/open-map", openmap.GetIndex)

	log.Fatal(http.ListenAndServe(commons.GetPort(), router))
}
