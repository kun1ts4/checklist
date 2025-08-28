package internal

import "time"

type Task struct {
	Id        string `gorm:"primary_key"`
	Text      string `gorm:"not null;type:text"`
	IsDone    bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
