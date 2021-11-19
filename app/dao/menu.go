package dao

import (
	"helloadmin/app/models"
)

func (d Dao) GetChildren(parentId uint) ([]*models.Menu, error) {
	menu := models.Menu{ParentId: parentId}
	return menu.Tree(d.engine)
}

func (d Dao) GetOptions() ([]map[string]interface{}, error) {
	menu := models.Menu{}
	return menu.Options(d.engine)
}

func (d Dao) Count(parentId uint) int64 {
	menu := models.Menu{ParentId: parentId}
	return menu.Count(d.engine)
}
