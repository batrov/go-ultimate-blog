package rcalc_statistics

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Batrov/go-ultimate-blog/commons"
	"gorm.io/gorm"
)

type FireCalcI interface {
	PostStatistics(params commons.RCalcStatistic) error
	GetStatistics() ([]commons.RCalcStatistic, error)
}

type FireCalc struct {
	DB *gorm.DB
}

func (r *FireCalc) PostStatistics(params commons.RCalcStatistic) (err error) {
	db := r.DB
	if db != nil {
		db.Table("rcalc_statistics").Create(&params)
		return
	}

	// check if file exists

	jsonFile, err := os.Open(commons.FireCalcJsonPath)
	if err != nil {
		// Try create file
		isFileNotFound := strings.Contains(err.Error(), "no such file or directory")
		if isFileNotFound {
			os.Create(commons.FireCalcJsonPath)
			fmt.Println("File created: ", commons.FireCalcJsonPath)
			jsonFile, err = os.Open(commons.FireCalcJsonPath)
			if err != nil {
				return
			}
		} else {
			err = commons.Error(err, "PS_00")
			return
		}
	}

	defer jsonFile.Close()

	byteValue, err := os.ReadFile(commons.FireCalcJsonPath)
	if err != nil {
		return
	}

	var statsData []commons.RCalcStatistic

	json.Unmarshal(byteValue, &statsData)

	statsData = append(statsData, params)

	newFile, err := json.MarshalIndent(statsData, "", "    ")
	if err != nil {
		err = commons.Error(err, "PS_01")
		return
	}

	err = os.WriteFile(commons.FireCalcJsonPath, newFile, 0644)
	if err != nil {
		err = commons.Error(err, "PS_02")
		return
	}

	return
}

func (r *FireCalc) GetStatistics() (params []commons.RCalcStatistic, err error) {
	db := r.DB
	if db != nil {
		db.Table("rcalc_statistics").Order("created_at DESC").Find(&params)
		return
	}

	jsonFile, err := os.Open(commons.FireCalcJsonPath)
	if err != nil {
		err = commons.Error(err, "GS_00")
		return
	}

	defer jsonFile.Close()

	byteValue, err := os.ReadFile(commons.FireCalcJsonPath)
	if err != nil {
		err = commons.Error(err, "GS_01")
		return
	}

	json.Unmarshal(byteValue, &params)

	return
}
