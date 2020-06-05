package goquiz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Batrov/go-ultimate-blog/commons"
)

// PostResetAnswer Service
func PostResetAnswer(createFile bool) (bool, error) {
	var (
		err         error
		surveyDatas commons.PostSurveyAnswerData
	)

	jsonFile, err := os.Open(commons.SurveyAnswerJsonPath)
	if err != nil {
		// Try create file
		if createFile {
			os.Create(commons.SurveyAnswerJsonPath)
			fmt.Println("File created: ", commons.SurveyAnswerJsonPath)
			PostResetAnswer(false)
		} else {
			return false, commons.Error(err, "PRA_00")
		}
	}

	defer jsonFile.Close()

	newFile, err := json.MarshalIndent(surveyDatas, "", "    ")
	if err != nil {
		return false, commons.Error(err, "PRA_01")
	}

	err = ioutil.WriteFile(commons.SurveyAnswerJsonPath, newFile, 0644)
	if err != nil {
		return false, commons.Error(err, "PRA_02")
	}

	return true, err
}
