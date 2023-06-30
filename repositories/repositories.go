package repositories

import (
	"os"

	"github.com/Batrov/go-ultimate-blog/repositories/contact"
	"github.com/Batrov/go-ultimate-blog/repositories/rcalc_statistics"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repositories struct {
	ContactRepo    contact.ContactI
	RCalcStatsRepo rcalc_statistics.FireCalcI
}

var repo Repositories

func Init() (err error) {
	var (
		db *gorm.DB
	)

	dbDriver := os.Getenv("DATABASE_DRIVER")
	dsn := os.Getenv("DATABASE_URL")
	if dbDriver == "POSTGRES" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else { // use sqlite by default
		if len(dsn) == 0 {
			dsn = "database1.db"
		}
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}
	repo = Repositories{
		ContactRepo: &contact.Contact{
			DB: db,
		},
		RCalcStatsRepo: &rcalc_statistics.FireCalc{
			DB: db,
		},
	}
	return err
}

func Get() Repositories {
	return repo
}
