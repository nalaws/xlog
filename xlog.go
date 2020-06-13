package xlog

import (
	"fmt"
	"os"
	"strings"
)

type Xlog struct {
	logSwitch bool   // 日志总开关  true: 输出日志, false: 不输出日志
	logLevel  Level  // 定义日志级别
	teeFile   bool   // 是否输出到文件
	path      string // 如果输出到文件,可以指定文件名。默认为 <application>.log
}

func NewXlog() *Xlog {
	return &Xlog{
		logSwitch: true,
		logLevel:  Trace,
		teeFile:   false,
	}
}

// 设置日志开关
func (x *Xlog) SetLogSwitch(b bool) {
	x.logSwitch = b
}

// 设置日志级别
func (x *Xlog) SetLogLevel(l Level) {
	x.logLevel = l
}

// 设置日志存储到文件
func (x *Xlog) SetTeeFile(b bool) error {
	x.teeFile = b
	if f == nil {
		if x.path == "" {
			pos := strings.LastIndex(os.Args[0], ".")
			x.path = fmt.Sprintf("%s.log", os.Args[0][:pos])
		}
		return openLogFile(x.path)
	}

	return nil
}
