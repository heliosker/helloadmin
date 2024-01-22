package api

type DepartmentResponseItem struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	ParentId  int    `json:"parentId"`
	Leader    string `json:"leader"`
	UpdateAt  string `json:"updateAt"`
	CreatedAt string `json:"createdAt"`
}

type DepartmentFindRequest struct {
	Name string `form:"name" binding:"max=50"`                              // 部门名称
	Page int    `form:"page" binding:"required,min=1" example:"1"`          // 分页
	Size int    `form:"size" binding:"required,min=1,max=100" example:"10"` // 页码
}

type DepartmentCreateRequest struct {
	Name     string `json:"name" binding:"required"` // 部门名称
	Sort     int    `json:"sort" binding:"required"` // 排序值，值越大，显示顺序越靠前
	Leader   string `json:"leader"`                  // 部门负责人
	ParentId int    `json:"parentId"`                // 上级部门
}

type DepartmentUpdateRequest struct {
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	ParentId int    `json:"parentId"`
	Leader   string `json:"leader"`
}

type DepartmentResponse struct {
	Items      []DepartmentResponseItem `json:"items"`
	Pagination `json:"pagination"`
}
