package home

import (
	"github.com/Batrov/go-ultimate-blog/commons"
	"github.com/Batrov/go-ultimate-blog/repositories"
)

type HomeI interface {
	PostContact(params commons.Contact) (err error)
}

type Home struct {
	repositories repositories.Repositories
}

func Init() (HomeI, error) {
	repo := repositories.Get()
	return &Home{
		repositories: repo,
	}, nil
}
