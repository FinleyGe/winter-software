package utility

import (
	"os"
)

// 文件夹不存在则创建
func Mkdir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0777)
	}
}
