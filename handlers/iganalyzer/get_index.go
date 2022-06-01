package iganalyzer

import (
	"fmt"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/ahmdrz/goinsta/v2"
	"github.com/julienschmidt/httprouter"
)

// GetIndex Handler
func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err error
	)

	r.ParseForm()

	insta := goinsta.New("uname", "passwd")
	fmt.Printf("%+v\n", insta)

	if err = insta.Login(); err != nil {
		fmt.Printf("%+v\n%+v\n", err, insta.Challenge)
	}

	fileName := "get_iganalyzer_index.template"

	t := commons.GetTemplate().Lookup(commons.TemplateLayoutName)

	templateData := commons.TemplateData{
		Title: "Instagram Analyzer",
		Contents: map[string]interface{}{
			"Content": fileName,
			"JS_File": "",
			"Version": commons.GetVersion(),
		},
	}

	err = t.Execute(w, templateData)
	if err != nil {
		commons.PrintErr(err, "GIH_00")
	}

}
