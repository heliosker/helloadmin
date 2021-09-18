package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"reflect"
)

var (
	Cfg      *ini.File
	AppDebug bool
)

func init() {
	cfg, err := ini.Load("./config/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println("RUN_MODE: ", cfg.Section("").Key("RUN_MODE").String())
	size := cfg.Section("app").Key("PAGE_SIZE").String()
	fmt.Println("PAGE_SIZE: ", size)

	fmt.Println(reflect.TypeOf(size))

	i, err := cfg.Section("app").Key("HTTP_PORT").Int()
	fmt.Println("HTTP_PORT: ", i)
	fmt.Println(reflect.TypeOf(i))

}

func Load()  {

}
