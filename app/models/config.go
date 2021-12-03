package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Config struct {
	Model
	Module string `json:"module"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Remark string `json:"remark"`
}

type ConfigRet struct {
	Module string `json:"module"`
	Remark string `json:"remark"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type ConfigStore struct {
	ConfigItems []configItem
}

func (i *ConfigStore) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &i.ConfigItems)
}

type configItem struct {
	Key   string `form:"key" binding:"required"`
	Value string `form:"value" binding:"required"`
}

func (Config) TableName() string {
	return "hi_config"
}

func (c Config) List(db *gorm.DB) ([]*Config, error) {
	var cfg []*Config
	if c.Module != "" {
		db = db.Where("`module` = ?", c.Module)
	}
	if err := db.Find(&cfg).Error; err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (c Config) Store(db *gorm.DB) (int64, error) {
	ret := db.Table(c.TableName()).Where("`key` = ?", c.Key).Update("value", c.Value)
	if err := ret.Error; err != nil {
		return ret.RowsAffected, err
	}
	return ret.RowsAffected, nil
}

func (c Config) Val(db *gorm.DB) (string, error) {
	var cfg Config
	if err := db.Where("`key` = ?", c.Key).Find(&cfg).Error; err != nil {
		return "", err
	}
	return cfg.Value, nil
}
