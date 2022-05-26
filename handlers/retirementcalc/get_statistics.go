package retirementcalc

import (
	"encoding/json"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/services"
	"github.com/julienschmidt/httprouter"
)

// GetStatistics Handler
func GetStatistics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	r.ParseForm()

	if r.FormValue("secret") != "hermawan" {
		return
	}

	data, err := services.GetService().RCalcService.GetStatistics()
	if err != nil {
		commons.PrintErr(err, "GS_00")
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
