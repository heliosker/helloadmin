package dao

import "helloadmin/app/models"

func (d *Dao) GetAuth(username, password string) (models.Auth, error) {
	auth := models.Auth{Username: username, Password: password}
	return auth.Get(d.engine)
}
