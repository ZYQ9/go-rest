package models

import "time"

// Quest Table structure
type Quest struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reward      int       `json:"reward"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Player Table structure
type Player struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Class     string    `json:"class"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
