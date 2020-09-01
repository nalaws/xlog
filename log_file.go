package xlog

import (
	"bufio"
	"os"
	"path"
	"sync"
)

// 日志文件结构
type XlogFile struct {
	f       *os.File
	mu      sync.Mutex
	preFile *XlogFile
}

// 检测文件
func IsFileExist(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, err
	}
	return true, nil
}

// 日志文件是否为空
func (x *XlogFile) IsEmpty() bool {
	return x.f == nil && x.preFile == nil
}

// 打开文件
func (x *XlogFile) openLogFile(pathName string) error {
	exist, err := IsFileExist(pathName)
	if !exist || err != nil {
		if err = os.MkdirAll(path.Dir(pathName), 0755); err != nil {
			return err
		}
		if x.f, err = os.Create(pathName); err != nil {
			return err
		}
		return nil
	}

	x.f, err = os.OpenFile(pathName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 关闭文件
func (x *XlogFile) closeLogFile() error {
	if x.preFile != nil {
		x.preFile.closeLogFile()
		x.preFile = nil
	}
	err := x.f.Sync()
	if err != nil {
		return err
	}
	err = x.f.Close()
	if err != nil {
		return err
	}
	x.f = nil
	return nil
}

// 写日志文件
func (x *XlogFile) Log2File(data []byte) {
	x.mu.Lock()
	defer x.mu.Unlock()
	// x.f != nil
	bufWriter := bufio.NewWriter(x.f)
	data = append(data, '\n')
	bufWriter.Write(data)
	bufWriter.Flush()
}
