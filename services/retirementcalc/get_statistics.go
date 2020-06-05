package retirementcalc

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Batrov/go-ultimate-blog/commons"
)

// GetStatistics Service
func GetStatistics() ([]commons.RetirementCalcPostStatisticsRequest, error) {
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

	return statsData, err
}
