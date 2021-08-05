package blog

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/julienschmidt/httprouter"
)

func GetIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		title    string
		fileName string
		err      error
	)
	lang := ps.ByName("lang")
	name := ps.ByName("name")

	switch name {
	case "golang-unit-test-tutorial-for-beginner-using-vscode":
		title = "Golang Unit Test Beginner Tutorial using VSCode"
		if lang == "en" {
			fileName = "blog-0-en.template"
		} else {
			fileName = "blog-0-en.template"
		}

	}

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: title,
		Contents: map[string]interface{}{
			"Content": fileName,
			"JS_File": "",
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GAH_00")
	}
}
