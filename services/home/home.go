package home

import (
	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/repositories"
)

type HomeI interface {
	PostContact(params commons.Contact) (err error)
}

type Home struct {
	repositories repositories.RepositoriesI
}

func Init() (HomeI, error) {
	repo, err := repositories.Init()
	return &Home{
		repositories: repo,
	}, err
}
