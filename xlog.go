package xlog

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Xlog struct {
	skip      int        // 日志堆栈层数, 默认为2, 如果进行二次封装,每封装一层值加1
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
	dirPos := strings.LastIndex(os.Args[0], string(os.PathSeparator))
	if dirPos < 0 {
		dirPos = 0
	} else {
		dirPos += 1
	}
	dir, _ := os.Executable()
	exPath := filepath.Dir(dir)
	if namePos < dirPos {
		namePos = len(os.Args[0])
	}
	return &Xlog{
		skip:      2,
		logSwitch: true,
		logLevel:  Trace,
		teeFile:   false,
		appDir:    exPath + string(os.PathSeparator),
		appName:   os.Args[0][dirPos:namePos],
	}
}

// 设置封装层数, 进行1次二次封装设置为 1
func (x *Xlog) SetSkip(skip int) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.skip = x.skip + skip
}

// 设置日志配置
func (x *Xlog) SetXlogConfig(conf *XlogConfig) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.logSwitch = conf.OnOff
	x.logLevel = conf.LogLevel
	x.teeFile = conf.OutFile
	x.fileConf = conf.XlogFile
}

// 设置日志文件配置
func (x *Xlog) SetXlogFileConfig(conf *LogFile) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.fileConf = *conf
	if x.fileConf.Dir == "" {
		x.fileConf.Dir = x.appDir
	}
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
		if x.fileConf.Dir != "" && !strings.HasSuffix(x.fileConf.Dir, string(os.PathSeparator)) {
			x.fileConf.Dir += string(os.PathSeparator)
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

// 获取APP名称作为tag
func (x *Xlog) AppName() string {
	return x.appName
}
