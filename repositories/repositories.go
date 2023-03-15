package repositories

import (
	"os"

	"github.com/Batrov/go-ultimate-blog/repositories/contact"
	"github.com/Batrov/go-ultimate-blog/repositories/rcalc_statistics"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repositories struct {
	ContactRepo    contact.ContactI
	RCalcStatsRepo rcalc_statistics.FireCalcI
}

var repo Repositories

func Init() (err error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	repo = Repositories{
		ContactRepo: &contact.Contact{
			DB: db,
		},
		RCalcStatsRepo: &rcalc_statistics.FireCalc{
			DB: db,
		},
	}
	return nil
}

func Get() Repositories {
	return repo
}
