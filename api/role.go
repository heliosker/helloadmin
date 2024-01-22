package api

type RoleCreateRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=50"  example:"test"`  // 角色名称
	Slug     string `json:"slug" binding:"required,min=1,max=50"  example:"test"`  // 唯一标识符
	Describe string `json:"describe" binding:"max=255" example:"this is describe"` // 角色描述
}

type RoleUpdateRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=50"  example:"test"`
	Slug     string `json:"slug" binding:"required,min=1,max=50"  example:"test"`
	Describe string `json:"describe" binding:"max=255" example:"this is describe"`
}

type RoleFindRequest struct {
	Name string `form:"name" binding:"max=50" example:"test"`               // 角色名称
	Slug string `form:"slug" binding:"max=50"  example:"test"`              // 唯一标识符
	Page int    `form:"page" binding:"required,min=1" example:"1"`          // 分页
	Size int    `form:"size" binding:"required,min=1,max=100" example:"10"` // 页码
}

type RoleResponse struct {
	Items      []RoleResponseItem `json:"items"`
	Pagination `json:"pagination"`
}

type RoleResponseItem struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Describe  string `json:"describe"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
