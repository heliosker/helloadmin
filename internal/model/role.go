package model

import (
	"time"
)

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Describe  string    `json:"describe"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Menus     []Menu    `gorm:"many2many:role_menu;"`
}

func (m *Role) TableName() string {
	return "role"
}
