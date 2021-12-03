package upload

import (
	"fmt"
	"helloadmin/pkg/utils"
	"path"
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
	t := time.Now()
	return fmt.Sprintf("%s/%d/%d/", "storage/uploads", t.Year(), int(t.Month()))
}
