package model

import (
	"time"
)

type Department struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(30);not null;default:'';comment:部门名称"`
	ParentId  int       `json:"parent_id" gorm:"type:int;default:0;comment:上级部门ID"`
	Sort      int       `json:"sort" gorm:"type:int; default:0;comment:排序值，值越大越靠前"`
	Leader    string    `json:"leader" gorm:"type:varchar(60);not null;default:'';comment:部门负责人"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"default:null;comment:创建于"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"default:null;comment:更新于"`
}

func (u *Department) TableName() string {
	return "department"
}
