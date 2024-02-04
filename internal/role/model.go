package role

import (
	"helloadmin/internal/menu"
	"time"
)

type Model struct {
	ID        uint         `json:"id" gorm:"primaryKey"`
	Name      string       `json:"name" gorm:"type:varchar(60);not null;default:'';comment:角色名称"`
	Describe  string       `json:"describe" gorm:"type:varchar(255);not null;default:'';comment:角色描述"`
	CreatedAt time.Time    `json:"created_at,omitempty" gorm:"default:null;comment:创建于"`
	UpdatedAt time.Time    `json:"updated_at,omitempty" gorm:"default:null;comment:更新于"`
	Menus     []menu.Model `json:"_" gorm:"many2many:role_menu;"`
}

func (m *Model) TableName() string {
	return "role"
}
