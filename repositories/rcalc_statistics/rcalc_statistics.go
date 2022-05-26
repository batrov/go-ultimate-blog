package rcalc_statistics

import (
	"github.com/Batrov/go-ultimate-blog/commons"
	"gorm.io/gorm"
)

type RetirementCalcI interface {
	PostStatistics(params commons.RCalcStatistic) error
	GetStatistics() ([]commons.RCalcStatistic, error)
}

type RetirementCalc struct {
	DB *gorm.DB
}

func (r *RetirementCalc) PostStatistics(params commons.RCalcStatistic) (err error) {
	db := r.DB
	db.Table("rcalc_statistics").Create(&params)
	return
}

func (r *RetirementCalc) GetStatistics() (params []commons.RCalcStatistic, err error) {
	db := r.DB
	db.Table("rcalc_statistics").Order("created_at DESC").Find(&params)
	return
}
