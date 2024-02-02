package operation_record

import "time"

type Model struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserId     string    `json:"userId" gorm:"type:varchar(64);not null;default:'';index:idx_user_id;unique;comment:账号唯一ID"`
	Operation  string    `json:"operation" gorm:"type:varchar(64);not null;default:'';comment:操作类型"`
	Path       string    `json:"path" gorm:"type:varchar(255);not null;default:'';comment:操作路径"`
	Method     string    `json:"method" gorm:"type:varchar(128);not null;default:'';comment:请求方式"`
	Ip         string    `json:"ip" gorm:"type:varchar(60);not null;default:'';comment:客户端IP"`
	HttpStatus int       `json:"httpStatus" gorm:"type:int;not null;default:0;comment:HTTP状态码"`
	Payload    string    `json:"payload" gorm:"type:text;comment:请求参数"`
	Response   string    `json:"response" gorm:"type:text;comment:响应结果"`
	CreatedAt  time.Time `json:"createdAt,omitempty" gorm:"default:null;comment:创建于"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" gorm:"default:null;comment:更新于"`
}

func (m *Model) TableName() string {
	return "operation_record"
}
