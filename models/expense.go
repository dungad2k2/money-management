package models
import "time"

type Expense struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount    float64   `json:"amount"`
	Reason    string    `json:"reason"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}