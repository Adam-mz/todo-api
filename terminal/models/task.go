package models

import "time"

type Task struct {
	ID         int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Task       string     `json:"task" gorm:"not null"`
	IsDone     bool       `json:"is_done" gorm:"default:false"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at" gorm:"default:null"`
	UserID     uint       `json:"user_id"`
}
