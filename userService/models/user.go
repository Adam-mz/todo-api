package models

import "time"

type Task struct {
	ID         int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Email      string     `json:"emailsk" gorm:"not null"`
	Password   string     `json:"Password"`
	Created_at time.Time  `json:"created_at"`
	Updated_at time.Time  `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at" gorm:"default:null"`
}
