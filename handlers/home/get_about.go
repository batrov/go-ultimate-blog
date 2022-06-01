package home

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/julienschmidt/httprouter"
)

// GetIndex Handler
func GetAbout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err error
	)

	r.ParseForm()

	fileName := "get_home_about.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "About Me",
		Contents: map[string]interface{}{
			"Content": fileName,
			"Version": commons.GetVersion(),
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GAH_00")
	}

}
