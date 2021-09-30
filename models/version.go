package models

import (
	"gorm.io/gorm"
)

type Version struct {
	gorm.Model
	Ver         string `json:"ver"`
	Description string `json:"description"`
}

func (Version) TableName() string {
	return "hi_version"
}

