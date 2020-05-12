package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetQuiz Handler
func GetQuiz(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// var (
	// 	data   interface{}
	// 	err    error
	// 	params map[string]interface{}
	// )

	// data, err = services.GetQuiz(params)
	// if err != nil {
	// 	commons.PrintErr(err, "GSH_00")
	// }

	// response := data.(commons.GetQuizResponse)

	// fileName := "get_survey.template"

	// t, err := template.New(commons.TemplateLayoutName).ParseFiles(
	// 	commons.GetTemplatePath(commons.TemplateLayoutName),
	// 	commons.GetTemplatePath(fileName),
	// )
	// if err != nil {
	// 	commons.PrintErr(err, "GSH_01")
	// }

	// templateData := commons.TemplateData{
	// 	Title: "Survey",
	// 	Contents: map[string]interface{}{
	// 		"Content": fileName,
	// 	},
	// }

	// err = t.Execute(w, templateData)
	// if err != nil {
	// 	commons.PrintErr(err, "GSH_02")
	// }
}
