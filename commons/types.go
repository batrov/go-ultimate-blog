package commons

type GetIndexResponse struct {
	FullName string
}

type SurveyAnswer struct {
	String string `json:"string"`
	Counts int64  `json:"counts"`
}

type SurveyData struct {
	ID       int64          `json:"id"`
	Category string         `json:"category"`
	Question string         `json:"question"`
	Answers  []SurveyAnswer `json:"answers"`
}

type GetSurveyData struct {
	SurveyDatas []SurveyData `json:"survey_data"`
}

type GetSurveyResponse struct {
	SurveyData         SurveyData
	SubmittedSurveyIDs []int64
}

type GetQuizResponse struct {
	FullName string
}

type TemplateData struct {
	Title    string
	Contents map[string]interface{}
}
