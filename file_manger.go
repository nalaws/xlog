package xlog

import (
	"sync"
)

type XlogFileManager struct {
	xlogs map[string]*XlogFile
	mu    sync.Mutex
}

var (
	gxlogs     *XlogFileManager
	gxlogslock sync.Mutex
)

// 获取日志文件管理单例
func Instance() *XlogFileManager {
	if gxlogs == nil {
		gxlogslock.Lock()
		if gxlogs == nil {
			gxlogs = &XlogFileManager{
				xlogs: make(map[string]*XlogFile),
			}
		}
		gxlogslock.Unlock()
	}

	return gxlogs
}

// 获取单个日志文件
func (xlm *XlogFileManager) XlogFile(path, tag string) *XlogFile {
	xlf := xlm.xlogs[tag]
	if xlf == nil {
		xlm.mu.Lock()
		if xlf == nil {
			xlf = &XlogFile{}
			xlf.openLogFile(path)
			xlm.xlogs[tag] = xlf
		}
		xlm.mu.Unlock()
	}
	return xlf
}

func (xlm *XlogFileManager) CloseAll() {
	for _, v := range xlm.xlogs {
		if !v.IsEmpty() {
			v.closeLogFile()
		}
	}
}
