package handlers

import (
	"net/http"

	"github.com/Batrov/go-quiz/commons"
	"github.com/Batrov/go-quiz/services"
	"github.com/julienschmidt/httprouter"
)

// GetIndex Handler
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		data   interface{}
		err    error
		params map[string]interface{}
	)

	r.ParseForm()

	params = map[string]interface{}{
		"full_name": r.FormValue("full_name"),
	}

	data, err = services.GetIndex(params)
	if err != nil {
		commons.PrintErr(err, "GIH_00")
	}

	response := data.(commons.GetIndexResponse)

	fileName := "get_index.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "Index",
		Contents: map[string]interface{}{
			"Content":  fileName,
			"FullName": response.FullName,
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GIH_02")
	}

}
