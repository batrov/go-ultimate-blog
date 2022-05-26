package repositories

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type RepositoriesI interface {
	ContactI
}

type Repositories struct {
	DB *gorm.DB
}

func Init() (RepositoriesI, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return &Repositories{
		DB: db,
	}, err
}
