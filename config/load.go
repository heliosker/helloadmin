package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func Load() *ini.File {
	dir, e := os.Getwd()
	fmt.Println(dir)
	fmt.Println(e)
	fmt.Println("-----------")
	app, err := ini.Load("./.env")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	//version = app.Section("").Key("APP_VERSION").String()
	return app
}
