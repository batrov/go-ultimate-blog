package home

import (
	"time"

	"github.com/Batrov/go-ultimate-blog/commons"
)

// PostContact Service
func (s *Home) PostContact(params commons.Contact) (err error) {
	params.CreatedAt = time.Now()
	return s.repositories.PostContact(params)
}
