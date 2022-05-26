package services

import (
	"github.com/Batrov/go-ultimate-blog/services/home"
)

type Services struct {
	HomeService home.HomeI
}

var service Services

func Init() error {
	homeService, err := home.Init()
	service = Services{
		HomeService: homeService,
	}
	return err
}

func GetService() Services {
	return service
}
