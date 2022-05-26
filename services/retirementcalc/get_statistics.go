package retirementcalc

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Batrov/go-ultimate-blog/commons"
)

// GetStatistics Service
func (s *RetirementCalc) GetStatistics() ([]commons.RetirementCalcPostStatisticsRequest, error) {
	var (
		err       error
		statsData []commons.RetirementCalcPostStatisticsRequest
	)

	jsonFile, err := os.Open(commons.RetirementCalcJsonPath)
	if err != nil {
		return statsData, commons.Error(err, "GS_00")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &statsData)

	// get from DB
	var newData []commons.RetirementCalcPostStatisticsRequest
	data, err := s.repositories.RCalcStatsRepo.GetStatistics()
	for _, v := range data {
		singleData := commons.RetirementCalcPostStatisticsRequest{
			Timestamp:      v.CreatedAt.String(),
			Age:            v.Age,
			Lifespan:       v.Lifespan,
			Income:         v.Income,
			Expenses:       v.Expenses,
			Inflation:      v.Inflation,
			Currency:       v.Currency,
			Raise:          v.Raise,
			AdvancedMode:   v.AdvancedMode,
			Investments:    v.Investments,
			Returns:        v.Returns,
			OtherExpenses:  v.OtherExpenses,
			CurrentSavings: v.CurrentSavings,
		}

		newData = append(newData, singleData)
	}

	return newData, err
}
