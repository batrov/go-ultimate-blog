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

	router.GET("/", commons.Middleware(home.GetIndex))
	router.GET("/about", commons.Middleware(home.GetAbout))

	// Go Quiz
	router.GET("/goquiz", commons.Middleware(goquiz.GetIndex))
	router.GET("/goquiz/survey", goquiz.GetSurvey)
	router.POST("/goquiz/survey", commons.Middleware(goquiz.PostSurvey))
	router.GET("/goquiz/quiz", commons.Middleware(goquiz.GetQuiz))
	router.POST("/goquiz/reset-answer", commons.Middleware(goquiz.PostResetAnswer))
	router.GET("/goquiz/fetch-answer", commons.Middleware(goquiz.GetFetchAnswer))

	// Retirement Calculator
	router.GET("/retirement-calculator", commons.Middleware(retirementcalc.GetIndex))
	router.POST("/retirement-calculator/statistics", commons.Middleware(retirementcalc.PostStatistics))
	router.GET("/retirement-calculator/statistics", commons.Middleware(retirementcalc.GetStatistics))

	// Insta Analyzer
	router.GET("/instagram-analyzer", commons.Middleware(iganalyzer.GetIndex))

	// Open Map
	router.GET("/open-map", commons.Middleware(openmap.GetIndex))

	log.Fatal(http.ListenAndServe(commons.GetPort(), router))
}
