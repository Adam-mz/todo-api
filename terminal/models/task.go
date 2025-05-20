package models

type Task struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Task   string `json:"task" gorm:"not null"`
	IsDone bool   `json:"is_done" gorm:"default:false"`
	UserID uint   `json:"user_id"`
}
