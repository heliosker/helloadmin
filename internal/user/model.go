package user

import (
	"time"

	"helloadmin/internal/department"
	"helloadmin/internal/role"

	"gorm.io/gorm"
)

type Model struct {
	ID         uint             `json:"id" gorm:"primaryKey"`
	UserId     string           `json:"userId" gorm:"type:varchar(64);not null;default:'';index:idx_user_id;unique;comment:账号唯一ID"`
	Nickname   string           `json:"nickname" gorm:"type:varchar(64);default:'';comment:昵称"`
	Password   string           `json:"password" gorm:"type:varchar(255);not null;comment:密码"`
	Email      string           `json:"email" gorm:"type:varchar(60);not null;default:'';comment:邮箱"`
	Salt       string           `json:"salt" gorm:"type:varchar(60);not null;default:'';comment:盐字段"`
	RoleId     uint             `json:"roleId" gorm:"type:int;not null;default:0;comment:角色ID"`
	DeptId     uint             `json:"deptId" gorm:"type:int;not null;default:0;comment:部门ID"`
	Role       role.Model       `json:"_" gorm:"foreignKey:RoleId"`
	Department department.Model `json:"department" gorm:"foreignKey:DeptId"`
	CreatedAt  time.Time        `json:"createdAt" gorm:"default:null;comment:创建于"`
	UpdatedAt  time.Time        `json:"updatedAt" gorm:"default:null;comment:更新于"`
	DeletedAt  gorm.DeletedAt   `json:"deletedAt" gorm:"default:null;comment:删除于"`
}

func (u *Model) TableName() string {
	return "user"
}
