package models

type Menu struct {
	Model
	ParentId  int    `json:"parent_id"`
	Order     int    `json:"order"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Uri       string `json:"uri"`
	Extension string `json:"extension"`
	IsShow    int8   `json:"is_show"`
}

func (Menu) TableName() string {
	return "hi_menus"
}
