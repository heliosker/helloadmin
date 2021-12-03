package dao

import (
	"helloadmin/app/models"
)

func (d Dao) GetConfig(module string) ([]models.ConfigRet, error) {
	cfg := models.Config{Module: module}
	list, err := cfg.List(d.engine)
	var ret = make([]models.ConfigRet, len(list))
	if err != nil {
		return ret, err
	}
	for k, v := range list {
		ret[k].Module = v.Module
		ret[k].Key = v.Key
		ret[k].Value = v.Value
		ret[k].Remark = v.Remark
	}
	return ret, nil
}

func (d Dao) StoreConfig(items models.ConfigStore) error {
	var cfg = models.Config{}
	for _, v := range items.ConfigItems {
		cfg.Key = v.Key
		cfg.Value = v.Value
		if _, e := cfg.Store(d.engine); e != nil {
			return e
		}
	}
	return nil
}

func (d Dao) GetValue(key string) (string, error) {
	var cfg = models.Config{Key: key}
	return cfg.Val(d.engine)
}
