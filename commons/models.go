package commons

import "time"

type Contact struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Subject   string    `json:"subject" db:"subject"`
	Message   string    `json:"message" db:"message"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
