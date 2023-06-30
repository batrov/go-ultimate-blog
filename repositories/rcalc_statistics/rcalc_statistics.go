package rcalc_statistics

import (
	"github.com/Batrov/go-ultimate-blog/commons"
	"gorm.io/gorm"
)

type FireCalcI interface {
	PostStatistics(params commons.RCalcStatistic) error
	GetStatistics() ([]commons.RCalcStatistic, error)
}

type FireCalc struct {
	DB *gorm.DB
}

func (r *FireCalc) PostStatistics(params commons.RCalcStatistic) (err error) {
	db := r.DB
	if db == nil {
		return
	}
	db.Table("rcalc_statistics").Create(&params)
	return
}

func (r *FireCalc) GetStatistics() (params []commons.RCalcStatistic, err error) {
	db := r.DB
	if db == nil {
		return
	}
	db.Table("rcalc_statistics").Order("created_at DESC").Find(&params)
	return
}
