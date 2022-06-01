package firecalc

import (
	"github.com/Batrov/go-ultimate-blog/commons"
)

// GetStatistics Service
func (s *FireCalc) GetStatistics() ([]commons.FireCalcPostStatisticsRequest, error) {
	var (
		err       error
		statsData []commons.FireCalcPostStatisticsRequest
	)

	// get from DB
	data, err := s.repositories.RCalcStatsRepo.GetStatistics()
	for _, v := range data {
		singleData := commons.FireCalcPostStatisticsRequest{
			Timestamp:      v.CreatedAt.String(),
			Age:            v.Age,
			Lifespan:       v.Lifespan,
			Income:         v.Income,
			Expenses:       v.Expenses,
			Inflation:      v.Inflation,
			Currency:       v.Currency,
			Raise:          v.Raise,
			AdvancedMode:   v.AdvancedMode,
			Investments:    v.Investments,
			Returns:        v.Returns,
			OtherExpenses:  v.OtherExpenses,
			CurrentSavings: v.CurrentSavings,
		}

		statsData = append(statsData, singleData)
	}

	return statsData, err
}
