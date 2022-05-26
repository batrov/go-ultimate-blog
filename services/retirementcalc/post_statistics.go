package retirementcalc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/Batrov/go-ultimate-blog/commons"
)

func (s *RetirementCalc) PostStatistics(params commons.RetirementCalcPostStatisticsRequest, createFile bool) (bool, error) {
	var (
		err       error
		statsData []commons.RetirementCalcPostStatisticsRequest
		now       = time.Now()
	)

	params.Timestamp = now.String()

	jsonFile, err := os.Open(commons.RetirementCalcJsonPath)
	if err != nil {
		// Try create file
		if createFile {
			os.Create(commons.RetirementCalcJsonPath)
			fmt.Println("File created: ", commons.RetirementCalcJsonPath)
			s.PostStatistics(params, false)
		} else {
			return false, commons.Error(err, "PS_00")
		}
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &statsData)

	statsData = append(statsData, params)

	newFile, err := json.MarshalIndent(statsData, "", "    ")
	if err != nil {
		return false, commons.Error(err, "PS_01")
	}

	err = ioutil.WriteFile(commons.RetirementCalcJsonPath, newFile, 0644)
	if err != nil {
		return false, commons.Error(err, "PS_02")
	}

	// insert to DB
	err = s.repositories.RCalcStatsRepo.PostStatistics(commons.RCalcStatistic{
		Age:            params.Age,
		Lifespan:       params.Lifespan,
		Income:         params.Income,
		Expenses:       params.Expenses,
		Inflation:      params.Inflation,
		Currency:       params.Currency,
		Raise:          params.Raise,
		AdvancedMode:   params.AdvancedMode,
		Investments:    params.Investments,
		Returns:        params.Returns,
		OtherExpenses:  params.OtherExpenses,
		CurrentSavings: params.CurrentSavings,
		CreatedAt:      now,
	})
	if err != nil {
		return false, commons.Error(err, "PS_03")
	}

	return true, err
}
