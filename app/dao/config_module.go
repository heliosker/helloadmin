package dao

import "helloadmin/app/models"

func (d *Dao) GetConfigModules() ([]*models.ConfigModule, error) {
	cm := models.ConfigModule{}
	return cm.List(d.engine)
}

func (d *Dao) CreateConfigModule(module string) error {
	cm := models.ConfigModule{
		Module: module,
	}
	return cm.Create(d.engine)
}

func (d *Dao) UpdateConfigModule(id uint, module string) error {
	cm := models.ConfigModule{
		Module: module,
	}
	cm.ID = id
	return cm.Update(d.engine)
}

func (d *Dao) DeleteConfigModule(id uint) error {
	cm := models.ConfigModule{}
	cm.ID = id
	return cm.Delete(d.engine)
}
