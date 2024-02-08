package login_record

import (
	"helloadmin/internal/api"
)

type CreateRequest struct {
	Ip           string `json:"ip" binding:"max=60"`
	Os           string `json:"os"`
	Email        string `json:"email"`
	Browser      string `json:"browser"`
	Platform     string `json:"platform"`
	UserName     string `json:"userName"`
	ErrorMessage string `json:"ErrorMessage"`
}

type FindRequest struct {
	Ip    string `form:"ip" binding:"max=60"`
	Email string `form:"email" binding:"max=50"`
	Page  int    `form:"page" binding:"required,min=1" example:"1"`
	Size  int    `form:"size" binding:"required,min=1,max=100" example:"10"`
}

type Item struct {
	Ip           string `json:"ip"`
	Os           string `json:"os"`
	Email        string `json:"email"`
	UserName     string `json:"userName"`
	Browser      string `json:"browser"`
	Platform     string `json:"platform"`
	ErrorMessage string `json:"ErrorMessage"`
	UpdatedAt    string `json:"updatedAt"`
	CreatedAt    string `json:"createdAt"`
}

type Response struct {
	Items          []Item `json:"items"`
	api.Pagination `json:"pagination"`
}
