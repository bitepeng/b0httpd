package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func WriteFile(path string, data []byte) {
	// 如果文件夹不存在，预先创建文件夹
	if lastSeparator := strings.LastIndex(path, "/"); lastSeparator != -1 {
		dirPath := path[:lastSeparator]
		if _, err := os.Stat(dirPath); err != nil && os.IsNotExist(err) {
			os.MkdirAll(dirPath, os.ModePerm)
		}
	}

	// 已存在的文件，不应该覆盖重写，可能在前端更改了配置文件等
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err2 := ioutil.WriteFile(path, data, os.ModePerm); err2 != nil {
			fmt.Printf("Write file failed: %s\n", path)
		}
	} else {
		fmt.Printf("File exist, skip: %s\n", path)
	}
}

func ReadFile(path string) string {
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Read file failed: %s\n", path)
	}
	return string(f)
}
