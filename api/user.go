package api

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"`
	Password string `json:"password" binding:"required" example:"123456"`
	Nickname string `json:"nickname" binding:"required,max=50" example:"admin"`
	RoleId   uint   `json:"roleId" example:"1"`
	DeptId   uint   `json:"deptId" example:"1"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresAt   string `json:"expiresAt"`
	TokenType   string `json:"tokenType"`
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
