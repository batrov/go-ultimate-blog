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
	ID             int64     `json:"id" db:"id"`
	Age            int64     `json:"age" db:"age"`
	Lifespan       int64     `json:"lifespan" db:"lifespan"`
	Income         float64   `json:"income" db:"income"`
	Expenses       float64   `json:"expenses" db:"expenses"`
	Inflation      float64   `json:"inflation" db:"inflation"`
	Currency       string    `json:"currency" db:"currency"`
	Raise          float64   `json:"raise" db:"raise"`
	AdvancedMode   bool      `json:"advanced_mode" db:"advanced_mode"`
	Investments    float64   `json:"investments" db:"investments"`
	Returns        float64   `json:"returns" db:"returns"`
	OtherExpenses  int64     `json:"other_expenses" db:"other_expenses"`
	CurrentSavings float64   `json:"current_savings" db:"current_savings"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}
