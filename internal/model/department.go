package model

import (
	"time"
)

type Department struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	ParentId  int       `json:"parent_id"`
	Sort      int       `json:"sort"`
	Leader    string    `json:"leader"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *Department) TableName() string {
	return "department"
}
