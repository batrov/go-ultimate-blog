package goquiz

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Batrov/go-ultimate-blog/commons"
)

// PostSurvey Service
func PostSurvey(params commons.PostSurveyRequest) (bool, error) {
	var (
		err         error
		surveyDatas commons.GetSurveyData
	)

	if params.Answer == "" {
		err = errors.New("Cannot send empty answer")
		return false, commons.Error(err, "PS_00.0")
	}

	params.Answer = strings.TrimSpace(params.Answer)
	params.Answer = strings.ToLower(params.Answer)

	jsonFile, err := os.Open(commons.SurveyJsonPath)
	if err != nil {
		return false, commons.Error(err, "PS_00")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &surveyDatas)

	for i, val := range surveyDatas.SurveyDatas {
		if val.ID == params.SurveyID {
			var sameAnswer bool
			for j, val := range val.Answers {
				if val.String == params.Answer {
					surveyDatas.SurveyDatas[i].Answers[j].Counts++
					sameAnswer = true
				}
			}
			if !sameAnswer {
				surveyDatas.SurveyDatas[i].Answers = append(surveyDatas.SurveyDatas[i].Answers, commons.SurveyAnswer{Counts: 1, String: params.Answer})
			}
			break
		}
	}

	newFile, err := json.MarshalIndent(surveyDatas, "", "    ")
	if err != nil {
		return false, commons.Error(err, "PS_01")
	}

	err = ioutil.WriteFile(commons.SurveyJsonPath, newFile, 0644)
	if err != nil {
		return false, commons.Error(err, "PS_02")
	}

	return true, err
}
