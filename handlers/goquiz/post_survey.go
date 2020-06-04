package goquiz

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/services/goquiz"
	"github.com/julienschmidt/httprouter"
)

// PostSurvey Handler
func PostSurvey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var postSurveyRequest commons.PostSurveyRequest

	err := json.NewDecoder(r.Body).Decode(&postSurveyRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok, err := goquiz.PostSurvey(postSurveyRequest, true)
	if err != nil {
		commons.PrintErr(err, "GSH_00")
	}

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}

}
