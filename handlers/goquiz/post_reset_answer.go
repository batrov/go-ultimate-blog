package goquiz

import (
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/services/goquiz"
	"github.com/julienschmidt/httprouter"
)

// PostResetAnswer Handler
func PostResetAnswer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ok, err := goquiz.PostResetAnswer(true)
	if err != nil {
		commons.PrintErr(err, "PRA_00")
	}

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}

}
