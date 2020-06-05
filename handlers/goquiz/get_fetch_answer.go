package goquiz

import (
	"encoding/json"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/services/goquiz"
	"github.com/julienschmidt/httprouter"
)

// GetFetchAnswer Handler
func GetFetchAnswer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data, err := goquiz.GetFetchAnswer()
	if err != nil {
		commons.PrintErr(err, "PRA_00")
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)

		marshalData, _ := json.Marshal(data)

		w.Write(marshalData)
	}

}
