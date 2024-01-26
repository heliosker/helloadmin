package model

import (
	"time"
)

type Menu struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Icon      string    `json:"icon"`
	Path      string    `json:"path"`
	Type      string    `json:"type"`
	Action    string    `json:"action"`
	ParentId  uint      `json:"parent_id"`
	Component string    `json:"component"`
	Sort      int       `json:"sort"`
	Visible   string    `json:"visible"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Menu) TableName() string {
	return "menu"
}
