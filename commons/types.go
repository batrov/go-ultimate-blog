package commons

type GetIndexResponse struct {
	FullName string
}

type SurveyAnswer struct {
	String string `json:"string"`
	Counts int64  `json:"counts"`
}

type SurveyData struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Question string `json:"question"`
}

type SurveyAnswerData struct {
	ID      int64          `json:"id"`
	Answers []SurveyAnswer `json:"answers"`
}

type GetSurveyData struct {
	SurveyDatas []SurveyData `json:"survey_data"`
}

type PostSurveyAnswerData struct {
	SurveyAnswerDatas []SurveyAnswerData `json:"survey_answer_data"`
}

type GetSurveyResponse struct {
	SurveyData         SurveyData
	SubmittedSurveyIDs []int64
}

type GetQuizResponse struct {
	FullName string
}

type PostSurveyRequest struct {
	SurveyID int64  `json:"survey_id"`
	Answer   string `json:"answer"`
}

type TemplateData struct {
	Title    string
	Contents map[string]interface{}
}

type PostCalculateRetirementCalcParams struct {
	Age int64
}

type PostCalculateRetirementCalcResponse struct {
}
