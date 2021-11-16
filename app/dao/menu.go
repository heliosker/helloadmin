package dao

import (
	"helloadmin/app/models"
	"helloadmin/pkg/app"
)

func (d Dao) GetTreeMenu(title string, page, size int) ([]*models.Menu, error) {
	menu := models.Menu{Title: title}
	offset := app.GetPageOffset(page, size)
	return menu.List(d.engine, offset, size)
}

func (d Dao) GetOptions() (map[uint]string, error) {
	menu := models.Menu{}
	return menu.Options(d.engine)
}
