package firecalc

import (
	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/repositories"
)

type FireCalcI interface {
	PostStatistics(params commons.FireCalcPostStatisticsRequest, createFile bool) (bool, error)
	GetStatistics() ([]commons.FireCalcPostStatisticsRequest, error)
}

type FireCalc struct {
	repositories repositories.Repositories
}

func Init() (FireCalcI, error) {
	repo := repositories.Get()
	return &FireCalc{
		repositories: repo,
	}, nil
}
