package models

import "time"

// Quest Table structure
type Quest struct {
	QID         uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Reward      int       `json:"reward"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Player Table structure
type Player struct {
	PID       uint      `json:"id" gorm:"primary_key"`
	Class     string    `json:"class"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Assigned Quests Table structure
type Assigned struct {
	ID        uint     `json:"id" gorm:"primary_key"`
	PlayerID  []Player `json:"player_id" gorm:"foreignkey:PID"`
	QuestID   []Quest  `json:"quest_id" gorm:"foreignkey:QID"`
	Completed bool     `json:"completed"`
}
