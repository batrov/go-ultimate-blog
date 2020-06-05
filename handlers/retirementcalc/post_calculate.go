package retirementcalc

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// PostCalculate Handler
func PostCalculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// var (
	// 	err error
	// )

	// r.ParseForm()

	// params := commons.PostCalculateRetirementCalcParams{
	// 	Age: 22,
	// }

	// data, err := retirementcalc.PostCalculate(params)
	// if err != nil {
	// 	commons.PrintErr(err, "GIH_00")
	// }

	// w.Header().Set("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// encoded, _ := json.Marshal(data)

	// w.Write(encoded)

}
