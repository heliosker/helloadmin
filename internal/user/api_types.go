package user

import "helloadmin/internal/api"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"` // 邮箱
	Password string `json:"password" binding:"required" example:"123456"`                  // 密码
	Nickname string `json:"nickname" binding:"required,max=50" example:"admin"`            // 员工名
	RoleId   uint   `json:"roleId" example:"1"`                                            // 角色ID
	DeptId   uint   `json:"deptId" example:"1"`                                            // 部门ID
}

type FindRequest struct {
	Page     int    `form:"page" binding:"required,min=1" example:"1"`  // 页码
	Size     int    `form:"size" binding:"required,min=1" example:"10"` // 每页条数
	Email    string `form:"email" example:"admin@helloadmin.com"`       // 邮箱
	Nickname string `form:"nickname" example:"admin"`                   // 员工昵称
	RoleId   uint   `form:"roleId" example:"1"`                         // 角色ID
	DeptId   uint   `form:"deptId" example:"1"`                         // 部门ID
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
	RoleId   uint   `json:"roleId" example:"1"` // 角色ID
	DeptId   uint   `json:"deptId" example:"1"` // 部门ID
}

type ProfileData struct {
	Email  string `json:"email" example:"admin@helloadmin.com"`
	UserId string `json:"userId" example:"1"` // 用户ID
	RoleId uint   `json:"roleId" example:"1"` // 角色ID
	DeptId uint   `json:"deptId" example:"1"` // 部门ID
	Role   struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"role"`
	Department struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"department"`
	Nickname  string `json:"nickname" example:"Hi admin"`
	CreatedAt string `json:"createdAt" example:"2023-12-27 19:01:00"`
	UpdatedAt string `json:"updatedAt" example:"2023-12-27 19:01:00"`
}

type Response struct {
	Items          []ProfileData `json:"items"`
	api.Pagination `json:"pagination"`
}
