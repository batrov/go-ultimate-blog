package openmap

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

	fileName := "get_openmap_index.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "Open Map",
		Contents: map[string]interface{}{
			"Content": fileName,
			"JS_File": "openmap/get_index.js",
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GIH_00")
	}

}
