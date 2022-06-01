package services

import (
	"github.com/Batrov/go-ultimate-blog/services/firecalc"
	"github.com/Batrov/go-ultimate-blog/services/home"
)

type Services struct {
	HomeService  home.HomeI
	RCalcService firecalc.FireCalcI
}

var service Services

func Init() (err error) {
	homeService, err := home.Init()
	if err != nil {
		return
	}
	rcalcService, err := firecalc.Init()
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
