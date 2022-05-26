package repositories

import "github.com/Batrov/go-ultimate-blog/commons"

type ContactI interface {
	PostContact(params commons.Contact) (err error)
}

func (r *Repositories) PostContact(params commons.Contact) (err error) {
	db := r.DB
	db.Create(&params)
	return nil
}
