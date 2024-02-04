package user

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"` // 邮箱
	Password string `json:"password" binding:"required" example:"123456"`                  // 密码
	Nickname string `json:"nickname" binding:"required,max=50" example:"admin"`            // 员工名
	RoleId   uint   `json:"roleId" example:"1"`                                            // 角色ID
	DeptId   uint   `json:"deptId" example:"1"`                                            // 部门ID
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"` // 邮箱
	Password string `json:"password" binding:"required" example:"123456"`                  // 密码
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"` // 访问令牌
	ExpiresAt   string `json:"expiresAt"`   // 过期日期
	TokenType   string `json:"tokenType"`   // 令牌类型
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"admin"`
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"`
}

type GetProfileResponseData struct {
	UserId    string `json:"userId"`
	Nickname  string `json:"nickname" example:"hello"`
	Email     string `json:"email" example:"admin@helloadmin.com"`
	CreatedAt string `json:"createdAt" example:"2023-12-27 19:01:00"`
}
