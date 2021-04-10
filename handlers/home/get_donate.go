package home

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/julienschmidt/httprouter"
)

// GetDonate Handler
func GetDonate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err error
	)

	r.ParseForm()

	fileName := "get_home_donate.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "Fund Me!",
		Contents: map[string]interface{}{
			"Content": fileName,
			"JS_File": "home/get_donate.js",
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GDH_00")
	}

}
