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

type RCalcStatistic struct {
	ID             int64     `db:"id"`
	Age            int64     `db:"age"`
	Lifespan       int64     `db:"lifespan"`
	Income         float64   `db:"income"`
	Expenses       float64   `db:"expenses"`
	Inflation      float64   `db:"inflation"`
	Currency       string    `db:"currency"`
	Raise          float64   `db:"raise"`
	AdvancedMode   bool      `db:"advanced_mode"`
	Investments    float64   `db:"investments"`
	Returns        float64   `db:"returns"`
	OtherExpenses  int64     `db:"other_expenses"`
	CurrentSavings float64   `db:"current_savings"`
	CreatedAt      time.Time `db:"created_at"`
}
