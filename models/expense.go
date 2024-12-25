package models
import "time"

type Expense struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Amount    float64   `json:"amount"`
	Reason    string    `json:"reason"`
	UserID    int	    `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt time.Time `json:"created_at"`
}
