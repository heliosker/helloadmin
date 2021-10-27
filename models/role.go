package models

type Role struct {
	Model
	Name     string `json:"name"`
	Describe string `json:"describe"`
}

func (Role) TableName() string {
	return "hi_roles"
}
