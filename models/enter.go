package models

import "time"

type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IDRequest struct {
	ID uint `uri:"id" form:"id" json:"id"`
}
