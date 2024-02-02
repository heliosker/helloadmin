package model

import (
	"time"
)

type Menu struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(128);not null;default:'';comment:菜单名称"`
	Title     string    `json:"title" gorm:"type:varchar(128);not null;default:'';comment:菜单标题"`
	Icon      string    `json:"icon" gorm:"type:varchar(128);not null;default:'';comment:菜单图标"`
	Path      string    `json:"path" gorm:"type:varchar(255);not null;default:'';comment:菜单路径"`
	Type      string    `json:"type" gorm:"type:char(1);not null;default:'';comment:菜单类型 目录D 菜单M 按钮B"`
	ParentId  uint      `json:"parent_id" gorm:"type:int;default:0;comment:上级菜单ID"`
	Component string    `json:"component" gorm:"type:varchar(255);not null;default:'';comment:组件路径"`
	Sort      int       `json:"sort" gorm:"type:int; default:0;comment:排序值，值越大越靠前"`
	Visible   string    `json:"visible" gorm:"type:char(1);not null;default:'';comment:是否可见，Y可见 N不可见"`
	CreatedAt time.Time `json:"created_at" gorm:"default:null;comment:创建于"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:null;comment:更新于"`
}

func (u *Menu) TableName() string {
	return "menu"
}
