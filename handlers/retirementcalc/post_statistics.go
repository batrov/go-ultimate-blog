package retirementcalc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/services"
	"github.com/julienschmidt/httprouter"
)

// PostStatistics Handler
func PostStatistics(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var postStatsRequest commons.RetirementCalcPostStatisticsRequest

	err := json.NewDecoder(r.Body).Decode(&postStatsRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok, err := services.GetService().RCalcService.PostStatistics(postStatsRequest, true)
	if err != nil {
		commons.PrintErr(err, "PS_00")
	}

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}

}
