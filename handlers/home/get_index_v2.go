package home

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/julienschmidt/httprouter"
)

// GetIndexV2 Handler
func GetIndexV2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err error
	)

	fileName := "get_home2_index.template"

	t := commons.GetTemplate().Lookup(fileName)

	templateData := commons.TemplateData{
		Contents: map[string]interface{}{
			"Version": commons.GetVersion(),
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GIH_02")
	}

}
