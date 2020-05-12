package services

import "github.com/Batrov/go-quiz/commons"

// GetIndex Service
func GetIndex(params map[string]interface{}) (interface{}, error) {
	var (
		err  error
		data commons.GetIndexResponse
	)

	data.FullName = params["full_name"].(string)

	return data, err
}
