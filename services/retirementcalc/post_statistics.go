package retirementcalc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Batrov/go-ultimate-blog/commons"
)

func PostStatistics(params commons.RetirementCalcPostStatisticsRequest, createFile bool) (bool, error) {
	var (
		err       error
		statsData []commons.RetirementCalcPostStatisticsRequest
	)

	jsonFile, err := os.Open(commons.RetirementCalcJsonPath)
	if err != nil {
		// Try create file
		if createFile {
			os.Create(commons.RetirementCalcJsonPath)
			fmt.Println("File created: ", commons.RetirementCalcJsonPath)
			PostStatistics(params, false)
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

	return true, err
}
