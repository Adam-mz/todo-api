package models

import "time"

type Task struct {
	ID         int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Task       string     `json:"task" gorm:"not null"`
	IsDone     bool       `json:"is_done" gorm:"default:false"`
	Created_at time.Time  `json:"created_at"`
	Updated_at time.Time  `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at" gorm:"default:null"`
}
