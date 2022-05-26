package services

import (
	"github.com/Batrov/go-ultimate-blog/repositories"
)

type ServicesI interface {
}

type Services struct {
	repositories repositories.RepositoriesI
}

func Init() (ServicesI, error) {
	repo, err := repositories.Init()
	return &Services{
		repositories: repo,
	}, err
}
