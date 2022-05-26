package contact

import (
	"github.com/Batrov/go-ultimate-blog/commons"
	"gorm.io/gorm"
)

type ContactI interface {
	PostContact(params commons.Contact) (err error)
}

type Contact struct {
	DB *gorm.DB
}

func (r *Contact) PostContact(params commons.Contact) (err error) {
	db := r.DB
	db.Create(&params)
	return nil
}
