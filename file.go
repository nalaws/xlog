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
func openLogFile(pathName string) error {
	exist, err := IsFileExist(pathName)
	if !exist || err != nil {
		if err = os.MkdirAll(path.Dir(pathName), 0755); err != nil {
			return err
		}
		if f, err = os.Create(pathName); err != nil {
			return err
		}
		return nil
	}

	f, err = os.OpenFile(pathName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	return err
}

// 关闭文件
func closeLogFile() error {
	return f.Close()
}
