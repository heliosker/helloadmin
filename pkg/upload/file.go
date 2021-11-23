package upload

import (
	"helloadmin/pkg/utils"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	filename := strings.TrimSuffix(name, ext)
	filename = utils.EncodeMD5(filename)
	return filename + ext
}


func GetFileExt(name string) string {
	return path.Ext(name)
}

func SavePath() string {
	year, month, _ := time.Now().Date()
	y := strconv.Itoa(year)
	return "storage/uploads/" + y + "/" + month.String() + "/"
}
