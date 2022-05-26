package services

import (
	"github.com/Batrov/go-ultimate-blog/services/home"
	"github.com/Batrov/go-ultimate-blog/services/retirementcalc"
)

type Services struct {
	HomeService  home.HomeI
	RCalcService retirementcalc.RetirementCalcI
}

var service Services

func Init() (err error) {
	homeService, err := home.Init()
	if err != nil {
		return
	}
	rcalcService, err := retirementcalc.Init()
	if err != nil {
		return
	}
	service = Services{
		HomeService:  homeService,
		RCalcService: rcalcService,
	}
	return
}

func GetService() Services {
	return service
}
