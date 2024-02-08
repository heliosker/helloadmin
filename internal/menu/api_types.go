package menu

type CreateRequest struct {
	Name      string `json:"name" binding:"max=128"`                 // 菜单名称
	Title     string `json:"title" binding:"required,max=128"`       // 菜单标题
	Icon      string `json:"icon" binding:"max=128"`                 // 菜单图标
	Path      string `json:"path" binding:"max=255"`                 // 菜单路径
	Type      string `json:"type" binding:"required,eq=D|eq=M|eq=B"` // 菜单类型 目录D 菜单M 按钮B
	ParentId  uint   `json:"parentId" binding:"min=0"`               // 上级菜单ID
	Component string `json:"component"`                              // 组件路径
	Sort      int    `json:"sort"`                                   // 排序值，值越大越靠前
	Visible   string `json:"visible" binding:"required,eq=Y|eq=N"`   // 是否可见，Y可见 N不可见
}

type FindRequest struct {
	Name    string `form:"name"`    // 菜单名称
	Visible string `form:"visible"` // 是否可见，Y可见 N不可见
}

type ResponseItem struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Icon      string         `json:"icon"`
	Path      string         `json:"path"`
	Type      string         `json:"type"`
	ParentId  uint           `json:"parentId"`
	Component string         `json:"component"`
	Sort      int            `json:"sort"`
	Visible   string         `json:"visible"`
	CreatedAt string         `json:"createdAt"`
	UpdatedAt string         `json:"updatedAt"`
	Children  []ResponseItem `json:"children"`
}

type OptionRequest struct {
	Type string `form:"type" binding:"required,eq=D|eq=M|eq=B"` // 菜单类型 目录D 菜单M 按钮B
}

type Option struct {
	Label string `json:"label"`
	Value uint   `json:"value"`
}

type Response struct {
	Items []ResponseItem `json:"items"`
}

type UpdateRequest struct {
	Name      string `json:"name"`      // 菜单名称
	Title     string `json:"title"`     // 菜单标题
	Icon      string `json:"icon"`      // 菜单图标
	Path      string `json:"path"`      // 菜单路径
	Type      string `json:"type"`      // 菜单类型 目录D 菜单M 按钮B
	ParentId  uint   `json:"parentId"`  // 上级菜单ID
	Component string `json:"component"` // 组件路径
	Sort      int    `json:"sort"`      // 排序值，值越大越靠前
	Visible   string `json:"visible"`   // 是否可见，Y可见 N不可见
}
