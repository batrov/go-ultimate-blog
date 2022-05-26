package retirementcalc

import (
	"time"

	"github.com/Batrov/go-ultimate-blog/commons"
)

func (s *RetirementCalc) PostStatistics(params commons.RetirementCalcPostStatisticsRequest, createFile bool) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	// insert to DB
	err = s.repositories.RCalcStatsRepo.PostStatistics(commons.RCalcStatistic{
		Age:            params.Age,
		Lifespan:       params.Lifespan,
		Income:         params.Income,
		Expenses:       params.Expenses,
		Inflation:      params.Inflation,
		Currency:       params.Currency,
		Raise:          params.Raise,
		AdvancedMode:   params.AdvancedMode,
		Investments:    params.Investments,
		Returns:        params.Returns,
		OtherExpenses:  params.OtherExpenses,
		CurrentSavings: params.CurrentSavings,
		CreatedAt:      now,
	})
	if err != nil {
		return false, commons.Error(err, "PS_00")
	}

	return true, err
}
