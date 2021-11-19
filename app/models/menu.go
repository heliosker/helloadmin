package models

import "gorm.io/gorm"

type Menu struct {
	Model
	ParentId  uint   `json:"parent_id"`
	Sort      int    `json:"sort"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Uri       string `json:"uri"`
	Extension string `json:"extension"`
	IsShow    int8   `json:"is_show"`
}

func (Menu) TableName() string {
	return "hi_menus"
}

type MenuTree struct {
	Menu
	Children []Menu
}

func (m Menu) Tree(db *gorm.DB) ([]*Menu, error) {
	var menu []*Menu
	var err error
	db = db.Where("parent_id = ?", m.ParentId)
	if err = db.Order("sort DESC").Find(&menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}

func (m Menu) Options(db *gorm.DB) ([]map[string]interface{}, error) {
	var menu []*Menu
	var err error
	if err = db.Order("sort DESC").Find(&menu).Error; err != nil {
		return nil, err
	}
	var options = make([]map[string]interface{}, 0, len(menu))
	for i := 0; i < len(menu); i++ {
		ops := map[string]interface{}{"key": menu[i].ID, "value": menu[i].Title}
		options = append(options, ops)
	}
	return options, nil
}

func (m Menu) Count(db *gorm.DB) int64 {
	var count int64
	db.Model(m).Where("parent_id = ?", m.ParentId).Count(&count)
	return count
}
