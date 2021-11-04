package models

type AdminUser struct {
	Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Status   int8   `json:"status"`
	Password string `json:"password"`
	RoleId   int64  `json:"role_id"`
}

func (AdminUser) TableName() string {
	return "hi_admin_users"
}

func AdminUserExist(email, password string) bool {
	var admin AdminUser
	DB.Select("id").Where(AdminUser{Email: email, Password: password}).First(&admin)
	if admin.ID > 0 {
		return true
	}
	return false
}
