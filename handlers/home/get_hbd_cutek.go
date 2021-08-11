package home

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/julienschmidt/httprouter"
)

// GetHBDCutek Handler
func GetHBDCutek(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err error
	)

	fileName := "hbd-cutek.template"

	t := commons.GetTemplate().Lookup(fileName)

	templateData := commons.TemplateData{}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GIH_02")
	}

}
