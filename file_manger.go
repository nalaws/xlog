package xlog

import (
	"fmt"
	"sync"
	"time"
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
func (xlm *XlogFileManager) XlogFile(xclf *LogFile, tag string) *XlogFile {
	t := time.Now()
	xlf := xlm.xlogs[tag]
	if xlf == nil {
		xlm.mu.Lock()
		if xlf == nil {
			xlf = &XlogFile{}
			err := xlf.openLogFile(xlm.xlogFilePath(t, xclf, tag))
			fmt.Println("info:", err)
			if err != nil {
				fmt.Println("error:", err)
				xlm.mu.Unlock()
				return nil
			}
			xlm.xlogs[tag] = xlf
		}
		xlm.mu.Unlock()
	}
	switch xclf.Split {
	case 1:
		if t.Unix()-xlf.timestamp > xclf.SplitInterval {
			xlf.mu.Lock()
			if t.Unix()-xlf.timestamp > xclf.SplitInterval {
				oldxlf := &XlogFile{
					path:      xlf.path,
					f:         xlf.f,
					timestamp: xlf.timestamp,
					total:     xlf.total,
					preFile:   xlf.preFile,
				}
				xlf.preFile = oldxlf
				xlf.openLogFile(xlm.xlogFilePath(t, xclf, tag))
			}
			xlf.mu.Unlock()
		}
	case 2:
		if xlf.total >= xclf.SplitMax {
			xlf.mu.Lock()
			if xlf.total >= xclf.SplitMax {
				oldxlf := &XlogFile{
					path:      xlf.path,
					f:         xlf.f,
					timestamp: xlf.timestamp,
					total:     xlf.total,
					preFile:   xlf.preFile,
				}
				xlf.preFile = oldxlf
				xlf.openLogFile(xlm.xlogFilePath(t, xclf, tag))
			}
			xlf.mu.Unlock()
		}
	default:
	}

	return xlf
}

// 关闭所有日志
func (xlm *XlogFileManager) CloseAll() {
	for _, v := range xlm.xlogs {
		if !v.IsEmpty() {
			v.closeLogFile()
		}
	}
}

// 获取日志文件路径
func (xlm *XlogFileManager) xlogFilePath(t time.Time, xclf *LogFile, tag string) string {
	name := tag
	switch xclf.Split {
	case 1: // 按时间拆分
		if xclf.SplitInterval >= 24*60*60 {
			name = fmt.Sprintf("%s-%s", tag, t.Format(dayLayout))
		} else if xclf.SplitInterval >= 1*60*60 {
			name = fmt.Sprintf("%s-%s", tag, t.Format(hourLayout))
		} else {
			name = fmt.Sprintf("%s-%s", tag, t.Format(timeLayout))
		}
	case 2:
		name = fmt.Sprintf("%s-%s", tag, t.Format(timeLayout))
	default:
	}
	if xclf.Prefix != "" && name != "" {
		name = "-" + name
	}
	return fmt.Sprintf("%s%s%s%s.log", xclf.Dir, xclf.Prefix, name, xclf.Suffix)
}
