package api

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"admin@helloadmin.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
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
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}