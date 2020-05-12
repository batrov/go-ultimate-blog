package handlers

import (
	"net/http"

	"github.com/Batrov/go-quiz/commons"
	"github.com/Batrov/go-quiz/services"
	"github.com/julienschmidt/httprouter"
)

// GetSurvey Handler
func GetSurvey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		data   interface{}
		err    error
		params map[string]interface{}
	)

	params = map[string]interface{}{
		"submittedSurvey": []int64{},
	}

	data, err = services.GetSurvey(params)
	if err != nil {
		commons.PrintErr(err, "GSH_00")
	}

	response := data.(commons.GetSurveyResponse)

	fileName := "get_survey.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "Survey",
		Contents: map[string]interface{}{
			"Content":  fileName,
			"Question": response.SurveyData.Question,
			"SurveyID": response.SurveyData.ID,
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GSH_02")
	}
}
