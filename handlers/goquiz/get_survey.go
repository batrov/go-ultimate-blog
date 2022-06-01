package goquiz

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/services/goquiz"
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

	data, err = goquiz.GetSurvey(params)
	if err != nil {
		commons.PrintErr(err, "GSH_00")
	}

	response := data.(commons.GetSurveyResponse)

	fileName := "get_goquiz_survey.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "Survey",
		Contents: map[string]interface{}{
			"Content":          fileName,
			"Question":         response.SurveyData.Question,
			"SurveyID":         response.SurveyData.ID,
			"SubmittedAnswers": "",
			"Version":          commons.GetVersion(),
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GSH_02")
	}
}
