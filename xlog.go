package xlog

import (
	"os"
	"strings"
	"sync"
)

type Xlog struct {
	logSwitch bool       // 日志总开关  true: 输出日志, false: 不输出日志
	logLevel  Level      // 定义日志级别
	teeFile   bool       // 是否输出到文件
	lock      sync.Mutex // 日志打印互斥锁
	fileConf  LogFile    // 日志文件配置
	appDir    string
	appName   string
}

func NewXlog() *Xlog {
	namePos := strings.LastIndex(os.Args[0], ".")
	dirPos := strings.LastIndex(os.Args[0], "\\") + 1

	return &Xlog{
		logSwitch: true,
		logLevel:  Trace,
		teeFile:   false,
		appDir:    os.Args[0][:dirPos],
		appName:   os.Args[0][dirPos:namePos],
	}
}

func (x *Xlog) SetXlogConfig(conf *XlogConfig) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.logSwitch = conf.OnOff
	x.logLevel = conf.LogLevel
	x.teeFile = conf.OutFile
	x.fileConf = conf.XlogFile
}

// 设置日志开关
func (x *Xlog) SetLogSwitch(b bool) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.logSwitch = b
}

// 设置日志级别
func (x *Xlog) SetLogLevel(l Level) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.logLevel = l
}

// 设置日志存储到文件
func (x *Xlog) SetTeeFile(b bool) error {
	x.lock.Lock()
	defer x.lock.Unlock()
	if x.teeFile {
		x.Close()
	}
	x.teeFile = b
	if x.teeFile {
		if x.fileConf.Dir == "" {
			x.fileConf.Dir = x.appDir
		}
		if !strings.HasSuffix(x.fileConf.Dir, "\\") {
			x.fileConf.Dir += "\\"
		}
		XlogFileTimerStart()
	}

	return nil
}

// 关闭日志
func (x *Xlog) Close() {
	if x.teeFile {
		XlogFileTimerStop()
		Instance().CloseAll()
	}
}
