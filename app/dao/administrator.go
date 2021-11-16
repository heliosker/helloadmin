package dao

import (
	"helloadmin/app/models"
	"helloadmin/pkg/app"
)

func (d Dao) GetAdministrator(email, username string, status, page, size int) ([]*models.AdminUser, error) {
	user := models.AdminUser{Email: email, Username: username, Status: status}
	return user.Pagination(d.engine, app.GetPageOffset(page, size), size)
}

func (d Dao) CountAdministrator(email, username string, status int) (int64, error) {
	user := models.AdminUser{Email: email, Username: username, Status: status}
	return user.Count(d.engine)
}
