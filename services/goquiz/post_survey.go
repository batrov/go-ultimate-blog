package goquiz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Batrov/go-ultimate-blog/commons"
)

// PostSurvey Service
func PostSurvey(params commons.PostSurveyRequest, createFile bool) (bool, error) {
	var (
		err         error
		surveyDatas commons.PostSurveyAnswerData
	)

	if params.Answer == "" {
		err = errors.New("Cannot send empty answer")
		return false, commons.Error(err, "PS_00.0")
	}

	params.Answer = strings.TrimSpace(params.Answer)
	params.Answer = strings.ToLower(params.Answer)

	jsonFile, err := os.Open(commons.SurveyAnswerJsonPath)
	if err != nil {
		// Try create file
		if createFile {
			os.Create(commons.SurveyAnswerJsonPath)
			fmt.Println("File created: ", commons.SurveyAnswerJsonPath)
			PostSurvey(params, false)
		} else {
			return false, commons.Error(err, "PS_00")
		}
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &surveyDatas)

	var hasData bool
	for i, val := range surveyDatas.SurveyAnswerDatas {
		if val.ID == params.SurveyID {
			hasData = true
			var sameAnswer bool
			for j, val := range val.Answers {
				if val.String == params.Answer {
					surveyDatas.SurveyAnswerDatas[i].Answers[j].Counts++
					sameAnswer = true
				}
			}
			if !sameAnswer {
				surveyDatas.SurveyAnswerDatas[i].Answers = append(surveyDatas.SurveyAnswerDatas[i].Answers, commons.SurveyAnswer{Counts: 1, String: params.Answer})
			}
			break
		}
	}

	if !hasData {
		surveyDatas.SurveyAnswerDatas = append(surveyDatas.SurveyAnswerDatas, commons.SurveyAnswerData{
			ID: params.SurveyID,
			Answers: []commons.SurveyAnswer{
				commons.SurveyAnswer{
					String: params.Answer,
					Counts: 1,
				},
			},
		})
	}

	newFile, err := json.MarshalIndent(surveyDatas, "", "    ")
	if err != nil {
		return false, commons.Error(err, "PS_01")
	}

	err = ioutil.WriteFile(commons.SurveyAnswerJsonPath, newFile, 0644)
	if err != nil {
		return false, commons.Error(err, "PS_02")
	}

	return true, err
}
