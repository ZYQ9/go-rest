package models

import "time"

type Quest struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reward      int       `json:"reward"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Player struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Class     string    `json:"title"`
	Name      string    `json:"description"`
	Level     int       `json:"reward"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
