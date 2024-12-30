package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Moneyspend float64 `json:"moneyspend"`
}
