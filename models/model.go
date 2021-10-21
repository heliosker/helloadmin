package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var (
		err                          error
		dbName, user, password, host string
	)

	//cfg, err := config.Load().GetSection("database")
	if err != nil {
		fmt.Println("Fail to get section 'database': %v", err)
	}

	//dbName = cfg.Key("NAME").String()
	//user = cfg.Key("USER").String()
	//password = cfg.Key("PASSWORD").String()
	//host = cfg.Key("HOST").String()
	//tablePrefix = cfg.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}
