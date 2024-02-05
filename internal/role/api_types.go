package role

import (
	"helloadmin/internal/api"
)

type CreateRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=50"  example:"test"`  // 角色名称
	Describe string `json:"describe" binding:"max=255" example:"this is describe"` // 角色描述
}

type UpdateRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=50"  example:"test"`
	Describe string `json:"describe" binding:"max=255" example:"this is describe"`
}

type FindRequest struct {
	Name string `form:"name" binding:"max=50" example:"test"`               // 角色名称
	Page int    `form:"page" binding:"required,min=1" example:"1"`          // 分页
	Size int    `form:"size" binding:"required,min=1,max=100" example:"10"` // 页码
}

type MenuRequest struct {
	MenuId []uint `json:"menuId" binding:"required,unique,dive,gt=0"` // 菜单ID
}

type Response struct {
	Items          []ResponseItem `json:"items"`
	api.Pagination `json:"pagination"`
}

type ResponseItem struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Describe  string `json:"describe"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	MenuId    []uint `json:"menuId"`
}
