package model

import (
	"errors"
	"time"
)

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title" gorm:"not null ; size:255"`
	Done        bool       `json:"done" gorm:"not null ; default:false"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

func (Todo) TableName() string { return "todos" }

var (
	ErrNotFound   = errors.New("todo not found")
	ErrEmptyTitle = errors.New("title must not be empty")
)
