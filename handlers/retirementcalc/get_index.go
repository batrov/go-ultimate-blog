package retirementcalc

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/julienschmidt/httprouter"
)

// GetIndex Handler
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err error
	)

	r.ParseForm()

	fileName := "get_retirementcalc_index.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "Retirement Calculator",
		Contents: map[string]interface{}{
			"Content": fileName,
			"JS_File": "retirementcalc/get_index.js",
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GIH_02")
	}

}
