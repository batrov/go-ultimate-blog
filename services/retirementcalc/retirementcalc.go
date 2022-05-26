package retirementcalc

import (
	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/repositories"
)

type RetirementCalcI interface {
	PostStatistics(params commons.RetirementCalcPostStatisticsRequest, createFile bool) (bool, error)
	GetStatistics() ([]commons.RetirementCalcPostStatisticsRequest, error)
}

type RetirementCalc struct {
	repositories repositories.Repositories
}

func Init() (RetirementCalcI, error) {
	repo := repositories.Get()
	return &RetirementCalc{
		repositories: repo,
	}, nil
}
