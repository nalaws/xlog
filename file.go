package xlog

import (
	"os"
	"path"
)

var (
	f *os.File
)

// 检测文件
func IsFileExist(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}
	return true, nil
}

// 打开文件
func OpenLogFile(pathName string) error {
	var err error
	if !IsFileExist(pathName) {
		if err = os.MkdirAll(path.Dir(pathName), 0755); err != nil {
			return err
		}
		if f, err = os.Create(pathName); err != nil {
			return err
		}
	}

	return nil
}

// 关闭文件
func CloseLogFile() error {
	return f.Close()
}
