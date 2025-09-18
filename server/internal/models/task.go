package models

import "time"

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
	Priority    int        `json:"priority"`
	Status      string     `json:"status"` // todo|doing|done
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
