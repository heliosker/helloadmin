package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint      `json:"id",gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletedAt struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at",gorm:"index"`
}

func init() {
	var (
		err                              error
		dbName, username, password, host string
	)
	dbName = os.Getenv("DB_NAME")
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	host = os.Getenv("HOST")
	//tablePrefix = cfg.Key("TABLE_PREFIX").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		dbName,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
}
