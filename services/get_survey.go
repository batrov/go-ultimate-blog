package services

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Batrov/go-quiz/commons"
)

// GetSurvey Service
func GetSurvey(params map[string]interface{}) (interface{}, error) {
	var (
		err             error
		surveyDatas     commons.GetSurveyData
		data            commons.GetSurveyResponse
		submittedSurvey []int64
	)

	submittedSurvey = params["submittedSurvey"].([]int64)

	jsonFile, err := os.Open("assets/json/survey.json")
	if err != nil {
		return nil, commons.Error(err, "GS_00")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &surveyDatas)

	for _, val := range surveyDatas.SurveyDatas {
		var alreadySubmitted bool
		for _, val2 := range submittedSurvey {
			if val2 == val.ID {
				alreadySubmitted = true
			}
		}

		if !alreadySubmitted {
			data.SurveyData = val
			break
		}
	}

	return data, err
}
