package login_record

import "time"

type Model struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Ip           string    `json:"ip" gorm:"type:varchar(60);not null;default:'';comment:客户端IP"`
	Os           string    `json:"os" gorm:"type:varchar(60);comment:操作系统"`
	Email        string    `json:"email" gorm:"type:varchar(60);not null;default:'';comment:登录邮箱"`
	Browser      string    `json:"browser" gorm:"type:varchar(60);not null;default:'';comment:浏览器"`
	Platform     string    `json:"platform" gorm:"type:varchar(60);comment:平台"`
	ErrorMessage string    `json:"errorMessage" gorm:"type:varchar(255);comment:错误信息"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"default:null;comment:创建于"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"default:null;comment:更新于"`
}

func (m *Model) TableName() string {
	return "login_record"
}
