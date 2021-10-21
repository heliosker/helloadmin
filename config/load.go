package config

import (
	"fmt"
	"os"
)

func Load() string {
	dir, e := os.Getwd()
	fmt.Println(dir)
	fmt.Println(e)
	fmt.Println("-----------")
	//app, err := ini.Load("./.env")
	//if err != nil {
	//	fmt.Printf("Fail to read file: %v", err)
	//	os.Exit(1)
	//}

	//version = app.Section("").Key("APP_VERSION").String()
	//return app
	return ""
}
