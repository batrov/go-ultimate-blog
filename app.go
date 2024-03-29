package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/handlers/blog"
	"github.com/Batrov/go-ultimate-blog/handlers/firecalc"
	"github.com/Batrov/go-ultimate-blog/handlers/goquiz"
	"github.com/Batrov/go-ultimate-blog/handlers/home"
	"github.com/Batrov/go-ultimate-blog/handlers/iganalyzer"
	"github.com/Batrov/go-ultimate-blog/handlers/openmap"
	"github.com/Batrov/go-ultimate-blog/repositories"
	"github.com/Batrov/go-ultimate-blog/services"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
)

func main() {
	initServer()
	fmt.Printf("System running on http://localhost%s\n", commons.GetPort())
	fmt.Printf("Version: %s\n", commons.GetVersion())
	initRouter()
}

func initServer() {
	err := repositories.Init()
	if err != nil {
		panic(err)
	}

	err = services.Init()
	if err != nil {
		panic(err)
	}
}

func initRouter() {
	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	router.GET("/", commons.Middleware(home.GetIndexV2))
	router.GET("/about", commons.Middleware(home.GetAbout))
	router.POST("/contact", commons.Middleware(home.PostContact))
	router.GET("/donate", commons.Middleware(home.GetDonate))

	// Go Quiz
	router.GET("/goquiz", commons.Middleware(goquiz.GetIndex))
	router.GET("/goquiz/survey", commons.Middleware(goquiz.GetSurvey))
	router.POST("/goquiz/survey", commons.Middleware(goquiz.PostSurvey))
	router.GET("/goquiz/quiz", commons.Middleware(goquiz.GetQuiz))
	router.POST("/goquiz/reset-answer", commons.Middleware(goquiz.PostResetAnswer))
	router.GET("/goquiz/fetch-answer", commons.Middleware(goquiz.GetFetchAnswer))

	// FIRE Calculator
	router.GET("/fire-calculator", commons.Middleware(firecalc.GetIndex))
	router.POST("/fire-calculator/statistics", commons.Middleware(firecalc.PostStatistics))
	router.GET("/fire-calculator/statistics", commons.Middleware(firecalc.GetStatistics))

	// Insta Analyzer
	router.GET("/instagram-analyzer", commons.Middleware(iganalyzer.GetIndex))

	// Open Map
	router.GET("/open-map", commons.Middleware(openmap.GetIndex))

	// Blogs
	router.GET("/blog/:lang/:name", commons.Middleware(blog.GetIndex))

	log.Fatal(http.ListenAndServe(commons.GetPort(), router))
}
