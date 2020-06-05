package goquiz

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Batrov/go-ultimate-blog/commons"
)

// GetFetchAnswer Service
func GetFetchAnswer() (commons.PostSurveyAnswerData, error) {
	var (
		err         error
		surveyDatas commons.PostSurveyAnswerData
	)

	jsonFile, err := os.Open(commons.SurveyAnswerJsonPath)
	if err != nil {
		return surveyDatas, commons.Error(err, "PS_00")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &surveyDatas)

	return surveyDatas, err
}
