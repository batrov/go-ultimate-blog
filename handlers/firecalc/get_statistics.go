package firecalc

import (
	"fmt"
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

		html := `<table style="border: 1px solid #dddddd;">
					<tr>
					<th>Age</th>
					<th>Lifespan</th>
					<th>Income</th>
					<th>Expenses</th>
					<th>Inflation</th>
					<th>Currency</th>
					<th>Raise %</th>
					<th>AdvancedMode</th>
					<th>Investments</th>
					<th>Returns</th>
					<th>OtherExpenses</th>
					<th>CurrentSavings</th>
					<th>Created At</th>
					</tr>`
		for _, v := range data {
			html += `<tr>`
			html += fmt.Sprintf("<td>%v</td>", v.Age)
			html += fmt.Sprintf("<td>%v</td>", v.Lifespan)
			html += fmt.Sprintf("<td>%.2f</td>", v.Income)
			html += fmt.Sprintf("<td>%.2f</td>", v.Expenses)
			html += fmt.Sprintf("<td>%.2f</td>", v.Inflation)
			html += fmt.Sprintf("<td>%v</td>", v.Currency)
			html += fmt.Sprintf("<td>%.2f</td>", v.Raise)
			html += fmt.Sprintf("<td>%v</td>", v.AdvancedMode)
			html += fmt.Sprintf("<td>%.2f</td>", v.Investments)
			html += fmt.Sprintf("<td>%.2f</td>", v.Returns)
			html += fmt.Sprintf("<td>%v</td>", v.OtherExpenses)
			html += fmt.Sprintf("<td>%.2f</td>", v.CurrentSavings)
			html += fmt.Sprintf("<td>%v</td>", v.Timestamp)
			html += `</tr>`

		}
		html += `</table>`

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, html)
	}

}
